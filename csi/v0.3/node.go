/*
Package csi is CSI driver interface for OSD
Copyright 2017 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package csi

import (
	"fmt"
	"os"

	"github.com/libopenstorage/openstorage/api"
	"github.com/libopenstorage/openstorage/pkg/options"
	"github.com/libopenstorage/openstorage/pkg/util"

	csi "github.com/libopenstorage/openstorage/csi/v0.3/spec"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OsdCsiServer) NodeGetInfo(
	ctx context.Context,
	req *csi.NodeGetInfoRequest,
) (*csi.NodeGetInfoResponse, error) {

	clus, err := s.cluster.Enumerate()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to Enumerate cluster: %s", err)
	}

	result := &csi.NodeGetInfoResponse{
		NodeId: clus.NodeId,
	}

	return result, nil
}

// NodeGetId is a CSI API which gets the PX NodeId for the local node
func (s *OsdCsiServer) NodeGetId(
	ctx context.Context,
	req *csi.NodeGetIdRequest,
) (*csi.NodeGetIdResponse, error) {
	clus, err := s.cluster.Enumerate()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to Enumerate cluster: %s", err)
	}

	result := &csi.NodeGetIdResponse{
		NodeId: clus.NodeId,
	}

	logrus.Infof("NodeId is %s", result.NodeId)

	return result, nil
}

// NodePublishVolume is a CSI API call which mounts the volume on the specified
// target path on the node.
//
// TODO: Support READ ONLY Mounts
//
func (s *OsdCsiServer) NodePublishVolume(
	ctx context.Context,
	req *csi.NodePublishVolumeRequest,
) (*csi.NodePublishVolumeResponse, error) {

	logrus.Debugf("NodePublishVolume req[%#v]", req)

	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume id must be provided")
	}
	if len(req.GetTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path must be provided")
	}
	if req.GetVolumeCapability() == nil || req.GetVolumeCapability().GetAccessMode() == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume access mode must be provided")
	}

	// Get volume information
	v, err := util.VolumeFromName(s.driver, req.GetVolumeId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Volume id %s not found: %s",
			req.GetVolumeId(),
			err.Error())
	}
	if s.driver.Type() != api.DriverType_DRIVER_TYPE_BLOCK &&
		req.GetVolumeCapability().GetBlock() != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Trying to attach as block a non block device")
	}

	// Gather volume attributes
	spec, _, _, err := s.specHandler.SpecFromOpts(req.GetVolumeAttributes())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Invalid volume attributes: %#v",
			req.GetVolumeAttributes())
	}

	// This seems weird as a way to change opts to map[string]string
	opts := make(map[string]string)
	if len(spec.GetPassphrase()) != 0 {
		opts[options.OptionsSecret] = spec.GetPassphrase()
	}

	// If this is for a block driver, first attach the volume
	var devicePath string
	if s.driver.Type() == api.DriverType_DRIVER_TYPE_BLOCK {
		if devicePath, err = s.driver.Attach(ctx, req.GetVolumeId(), opts); err != nil {
			return nil, status.Errorf(
				codes.Internal,
				"Unable to attach volume: %s",
				err.Error())
		}
	}

	if req.GetVolumeCapability().GetBlock() != nil {
		// As block create a sym link to the attached location
		err = os.Symlink(devicePath, req.GetTargetPath())
		if err != nil {
			detachErr := s.driver.Detach(ctx, v.GetId(), opts)
			if detachErr != nil {
				logrus.Errorf("Unable to detach volume %s: %s",
					v.GetId(),
					detachErr.Error())
			}
			return nil, status.Errorf(
				codes.Internal,
				"Failed to create symlink %s -> %s: %v",
				req.GetTargetPath(),
				devicePath,
				err)
		}
	} else {
		// Verify target location is an existing directory
		if err := verifyTargetLocation(req.GetTargetPath()); err != nil {
			return nil, status.Errorf(
				codes.Aborted,
				"Failed to use target location %s: %s",
				req.GetTargetPath(),
				err.Error())
		}

		// Mount volume onto the path
		if err := s.driver.Mount(ctx, req.GetVolumeId(), req.GetTargetPath(), nil); err != nil {
			// Detach on error
			detachErr := s.driver.Detach(ctx, v.GetId(), opts)
			if detachErr != nil {
				logrus.Errorf("Unable to detach volume %s: %s",
					v.GetId(),
					detachErr.Error())
			}
			return nil, status.Errorf(
				codes.Internal,
				"Unable to mount volume %s onto %s: %s",
				req.GetVolumeId(),
				req.GetTargetPath(),
				err.Error())
		}
	}

	logrus.Infof("Volume %s mounted on %s",
		req.GetVolumeId(),
		req.GetTargetPath())

	return &csi.NodePublishVolumeResponse{}, nil
}

// NodeUnpublishVolume is a CSI API call which unmounts the volume.
func (s *OsdCsiServer) NodeUnpublishVolume(
	ctx context.Context,
	req *csi.NodeUnpublishVolumeRequest,
) (*csi.NodeUnpublishVolumeResponse, error) {

	logrus.Debugf("NodeUnPublishVolume req[%#v]", req)

	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume id must be provided")
	}
	if len(req.GetTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path must be provided")
	}

	// Get volume information
	_, err := util.VolumeFromName(s.driver, req.GetVolumeId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Volume id %s not found: %s",
			req.GetVolumeId(),
			err.Error())
	}

	// Get information about the target since the request does not
	// tell us if it is for block or mount point.
	// https://github.com/container-storage-interface/spec/issues/285
	fileInfo, err := os.Lstat(req.GetTargetPath())
	if err != nil && os.IsNotExist(err) {
		// For idempotency, return that there is nothing to unmount
		logrus.Infof("NodeUnpublishVolume on target path %s but it does "+
			"not exist, returning there is nothing to do", req.GetTargetPath())
		return &csi.NodeUnpublishVolumeResponse{}, nil
	} else if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Unknown error while verifying target location %s: %s",
			req.GetTargetPath(),
			err.Error())
	}

	// Check if it is block or not
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		// If block, we just need to remove the link.
		os.Remove(req.GetTargetPath())
	} else {
		if !fileInfo.IsDir() {
			return nil, status.Errorf(
				codes.NotFound,
				"Target location %s is not a directory", req.GetTargetPath())
		}

		// Mount volume onto the path
		if err = s.driver.Unmount(ctx, req.GetVolumeId(), req.GetTargetPath(), nil); err != nil {
			return nil, status.Errorf(
				codes.Internal,
				"Unable to unmount volume %s onto %s: %s",
				req.GetVolumeId(),
				req.GetTargetPath(),
				err.Error())
		}
	}

	if s.driver.Type() == api.DriverType_DRIVER_TYPE_BLOCK {
		if err = s.driver.Detach(ctx, req.GetVolumeId(), nil); err != nil {
			return nil, status.Errorf(
				codes.Internal,
				"Unable to detach volume: %s",
				err.Error())
		}
	}

	logrus.Infof("Volume %s unmounted", req.GetVolumeId())

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

// NodeGetCapabilities is a CSI API function which seems to be setup for
// future patches
func (s *OsdCsiServer) NodeGetCapabilities(
	ctx context.Context,
	req *csi.NodeGetCapabilitiesRequest,
) (*csi.NodeGetCapabilitiesResponse, error) {

	logrus.Debugf("NodeGetCapabilities req[%#v]", req)

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_UNKNOWN,
					},
				},
			},
		},
	}, nil
}

func verifyTargetLocation(targetPath string) error {
	fileInfo, err := os.Lstat(targetPath)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("Target location %s does not exist", targetPath)
	} else if err != nil {
		return fmt.Errorf(
			"Unknown error while verifying target location %s: %s",
			targetPath,
			err.Error())
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("Target location %s is not a directory", targetPath)
	}

	return nil
}
