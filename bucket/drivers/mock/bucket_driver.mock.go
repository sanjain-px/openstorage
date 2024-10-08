// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/bucket (interfaces: BucketDriver)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	bucket "github.com/libopenstorage/openstorage/bucket"
	reflect "reflect"
)

// MockBucketDriver is a mock of BucketDriver interface
type MockBucketDriver struct {
	ctrl     *gomock.Controller
	recorder *MockBucketDriverMockRecorder
}

// MockBucketDriverMockRecorder is the mock recorder for MockBucketDriver
type MockBucketDriverMockRecorder struct {
	mock *MockBucketDriver
}

// NewMockBucketDriver creates a new mock instance
func NewMockBucketDriver(ctrl *gomock.Controller) *MockBucketDriver {
	mock := &MockBucketDriver{ctrl: ctrl}
	mock.recorder = &MockBucketDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBucketDriver) EXPECT() *MockBucketDriverMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method
func (m *MockBucketDriver) CreateBucket(arg0, arg1, arg2 string, arg3 api.AnonymousBucketAccessMode) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockBucketDriverMockRecorder) CreateBucket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockBucketDriver)(nil).CreateBucket), arg0, arg1, arg2, arg3)
}

// DeleteBucket mocks base method
func (m *MockBucketDriver) DeleteBucket(arg0, arg1, arg2 string, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket
func (mr *MockBucketDriverMockRecorder) DeleteBucket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockBucketDriver)(nil).DeleteBucket), arg0, arg1, arg2, arg3)
}

// GrantBucketAccess mocks base method
func (m *MockBucketDriver) GrantBucketAccess(arg0, arg1, arg2 string) (string, *bucket.BucketAccessCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantBucketAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*bucket.BucketAccessCredentials)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GrantBucketAccess indicates an expected call of GrantBucketAccess
func (mr *MockBucketDriverMockRecorder) GrantBucketAccess(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantBucketAccess", reflect.TypeOf((*MockBucketDriver)(nil).GrantBucketAccess), arg0, arg1, arg2)
}

// RevokeBucketAccess mocks base method
func (m *MockBucketDriver) RevokeBucketAccess(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeBucketAccess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeBucketAccess indicates an expected call of RevokeBucketAccess
func (mr *MockBucketDriverMockRecorder) RevokeBucketAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeBucketAccess", reflect.TypeOf((*MockBucketDriver)(nil).RevokeBucketAccess), arg0, arg1)
}

// Start mocks base method
func (m *MockBucketDriver) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockBucketDriverMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBucketDriver)(nil).Start))
}

// String mocks base method
func (m *MockBucketDriver) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockBucketDriverMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockBucketDriver)(nil).String))
}
