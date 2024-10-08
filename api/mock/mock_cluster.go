// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStorageClusterServer,OpenStorageClusterClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockOpenStorageClusterServer is a mock of OpenStorageClusterServer interface
type MockOpenStorageClusterServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageClusterServerMockRecorder
}

// MockOpenStorageClusterServerMockRecorder is the mock recorder for MockOpenStorageClusterServer
type MockOpenStorageClusterServerMockRecorder struct {
	mock *MockOpenStorageClusterServer
}

// NewMockOpenStorageClusterServer creates a new mock instance
func NewMockOpenStorageClusterServer(ctrl *gomock.Controller) *MockOpenStorageClusterServer {
	mock := &MockOpenStorageClusterServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageClusterServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageClusterServer) EXPECT() *MockOpenStorageClusterServerMockRecorder {
	return m.recorder
}

// InspectCurrent mocks base method
func (m *MockOpenStorageClusterServer) InspectCurrent(arg0 context.Context, arg1 *api.SdkClusterInspectCurrentRequest) (*api.SdkClusterInspectCurrentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectCurrent", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkClusterInspectCurrentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectCurrent indicates an expected call of InspectCurrent
func (mr *MockOpenStorageClusterServerMockRecorder) InspectCurrent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectCurrent", reflect.TypeOf((*MockOpenStorageClusterServer)(nil).InspectCurrent), arg0, arg1)
}

// MockOpenStorageClusterClient is a mock of OpenStorageClusterClient interface
type MockOpenStorageClusterClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageClusterClientMockRecorder
}

// MockOpenStorageClusterClientMockRecorder is the mock recorder for MockOpenStorageClusterClient
type MockOpenStorageClusterClientMockRecorder struct {
	mock *MockOpenStorageClusterClient
}

// NewMockOpenStorageClusterClient creates a new mock instance
func NewMockOpenStorageClusterClient(ctrl *gomock.Controller) *MockOpenStorageClusterClient {
	mock := &MockOpenStorageClusterClient{ctrl: ctrl}
	mock.recorder = &MockOpenStorageClusterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageClusterClient) EXPECT() *MockOpenStorageClusterClientMockRecorder {
	return m.recorder
}

// InspectCurrent mocks base method
func (m *MockOpenStorageClusterClient) InspectCurrent(arg0 context.Context, arg1 *api.SdkClusterInspectCurrentRequest, arg2 ...grpc.CallOption) (*api.SdkClusterInspectCurrentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InspectCurrent", varargs...)
	ret0, _ := ret[0].(*api.SdkClusterInspectCurrentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectCurrent indicates an expected call of InspectCurrent
func (mr *MockOpenStorageClusterClientMockRecorder) InspectCurrent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectCurrent", reflect.TypeOf((*MockOpenStorageClusterClient)(nil).InspectCurrent), varargs...)
}
