// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStorageWatchServer,OpenStorageWatchClient,OpenStorageWatch_WatchClient,OpenStorageWatch_WatchServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockOpenStorageWatchServer is a mock of OpenStorageWatchServer interface
type MockOpenStorageWatchServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageWatchServerMockRecorder
}

// MockOpenStorageWatchServerMockRecorder is the mock recorder for MockOpenStorageWatchServer
type MockOpenStorageWatchServerMockRecorder struct {
	mock *MockOpenStorageWatchServer
}

// NewMockOpenStorageWatchServer creates a new mock instance
func NewMockOpenStorageWatchServer(ctrl *gomock.Controller) *MockOpenStorageWatchServer {
	mock := &MockOpenStorageWatchServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageWatchServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageWatchServer) EXPECT() *MockOpenStorageWatchServerMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockOpenStorageWatchServer) Watch(arg0 *api.SdkWatchRequest, arg1 api.OpenStorageWatch_WatchServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch
func (mr *MockOpenStorageWatchServerMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockOpenStorageWatchServer)(nil).Watch), arg0, arg1)
}

// MockOpenStorageWatchClient is a mock of OpenStorageWatchClient interface
type MockOpenStorageWatchClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageWatchClientMockRecorder
}

// MockOpenStorageWatchClientMockRecorder is the mock recorder for MockOpenStorageWatchClient
type MockOpenStorageWatchClientMockRecorder struct {
	mock *MockOpenStorageWatchClient
}

// NewMockOpenStorageWatchClient creates a new mock instance
func NewMockOpenStorageWatchClient(ctrl *gomock.Controller) *MockOpenStorageWatchClient {
	mock := &MockOpenStorageWatchClient{ctrl: ctrl}
	mock.recorder = &MockOpenStorageWatchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageWatchClient) EXPECT() *MockOpenStorageWatchClientMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockOpenStorageWatchClient) Watch(arg0 context.Context, arg1 *api.SdkWatchRequest, arg2 ...grpc.CallOption) (api.OpenStorageWatch_WatchClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(api.OpenStorageWatch_WatchClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockOpenStorageWatchClientMockRecorder) Watch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockOpenStorageWatchClient)(nil).Watch), varargs...)
}

// MockOpenStorageWatch_WatchClient is a mock of OpenStorageWatch_WatchClient interface
type MockOpenStorageWatch_WatchClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageWatch_WatchClientMockRecorder
}

// MockOpenStorageWatch_WatchClientMockRecorder is the mock recorder for MockOpenStorageWatch_WatchClient
type MockOpenStorageWatch_WatchClientMockRecorder struct {
	mock *MockOpenStorageWatch_WatchClient
}

// NewMockOpenStorageWatch_WatchClient creates a new mock instance
func NewMockOpenStorageWatch_WatchClient(ctrl *gomock.Controller) *MockOpenStorageWatch_WatchClient {
	mock := &MockOpenStorageWatch_WatchClient{ctrl: ctrl}
	mock.recorder = &MockOpenStorageWatch_WatchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageWatch_WatchClient) EXPECT() *MockOpenStorageWatch_WatchClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockOpenStorageWatch_WatchClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockOpenStorageWatch_WatchClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockOpenStorageWatch_WatchClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockOpenStorageWatch_WatchClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockOpenStorageWatch_WatchClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOpenStorageWatch_WatchClient)(nil).Context))
}

// Header mocks base method
func (m *MockOpenStorageWatch_WatchClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockOpenStorageWatch_WatchClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockOpenStorageWatch_WatchClient)(nil).Header))
}

// Recv mocks base method
func (m *MockOpenStorageWatch_WatchClient) Recv() (*api.SdkWatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*api.SdkWatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockOpenStorageWatch_WatchClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockOpenStorageWatch_WatchClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockOpenStorageWatch_WatchClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockOpenStorageWatch_WatchClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOpenStorageWatch_WatchClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockOpenStorageWatch_WatchClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockOpenStorageWatch_WatchClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOpenStorageWatch_WatchClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockOpenStorageWatch_WatchClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockOpenStorageWatch_WatchClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockOpenStorageWatch_WatchClient)(nil).Trailer))
}

// MockOpenStorageWatch_WatchServer is a mock of OpenStorageWatch_WatchServer interface
type MockOpenStorageWatch_WatchServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageWatch_WatchServerMockRecorder
}

// MockOpenStorageWatch_WatchServerMockRecorder is the mock recorder for MockOpenStorageWatch_WatchServer
type MockOpenStorageWatch_WatchServerMockRecorder struct {
	mock *MockOpenStorageWatch_WatchServer
}

// NewMockOpenStorageWatch_WatchServer creates a new mock instance
func NewMockOpenStorageWatch_WatchServer(ctrl *gomock.Controller) *MockOpenStorageWatch_WatchServer {
	mock := &MockOpenStorageWatch_WatchServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageWatch_WatchServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageWatch_WatchServer) EXPECT() *MockOpenStorageWatch_WatchServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockOpenStorageWatch_WatchServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockOpenStorageWatch_WatchServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOpenStorageWatch_WatchServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockOpenStorageWatch_WatchServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockOpenStorageWatch_WatchServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOpenStorageWatch_WatchServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockOpenStorageWatch_WatchServer) Send(arg0 *api.SdkWatchResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockOpenStorageWatch_WatchServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOpenStorageWatch_WatchServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockOpenStorageWatch_WatchServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockOpenStorageWatch_WatchServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockOpenStorageWatch_WatchServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockOpenStorageWatch_WatchServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockOpenStorageWatch_WatchServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOpenStorageWatch_WatchServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockOpenStorageWatch_WatchServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockOpenStorageWatch_WatchServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockOpenStorageWatch_WatchServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockOpenStorageWatch_WatchServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockOpenStorageWatch_WatchServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockOpenStorageWatch_WatchServer)(nil).SetTrailer), arg0)
}
