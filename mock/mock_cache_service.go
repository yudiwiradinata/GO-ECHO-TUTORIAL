// Code generated by MockGen. DO NOT EDIT.
// Source: GO-ECHO-TUTORIAL/service (interfaces: CacheServiceInterface)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCacheServiceInterface is a mock of CacheServiceInterface interface
type MockCacheServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCacheServiceInterfaceMockRecorder
}

// MockCacheServiceInterfaceMockRecorder is the mock recorder for MockCacheServiceInterface
type MockCacheServiceInterfaceMockRecorder struct {
	mock *MockCacheServiceInterface
}

// NewMockCacheServiceInterface creates a new mock instance
func NewMockCacheServiceInterface(ctrl *gomock.Controller) *MockCacheServiceInterface {
	mock := &MockCacheServiceInterface{ctrl: ctrl}
	mock.recorder = &MockCacheServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCacheServiceInterface) EXPECT() *MockCacheServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateCache mocks base method
func (m *MockCacheServiceInterface) CreateCache(arg0 string, arg1 interface{}, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCache", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCache indicates an expected call of CreateCache
func (mr *MockCacheServiceInterfaceMockRecorder) CreateCache(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCache", reflect.TypeOf((*MockCacheServiceInterface)(nil).CreateCache), arg0, arg1, arg2)
}

// DeleteCache mocks base method
func (m *MockCacheServiceInterface) DeleteCache(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCache", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCache indicates an expected call of DeleteCache
func (mr *MockCacheServiceInterfaceMockRecorder) DeleteCache(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCache", reflect.TypeOf((*MockCacheServiceInterface)(nil).DeleteCache), arg0)
}

// FindCache mocks base method
func (m *MockCacheServiceInterface) FindCache(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCache", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindCache indicates an expected call of FindCache
func (mr *MockCacheServiceInterfaceMockRecorder) FindCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCache", reflect.TypeOf((*MockCacheServiceInterface)(nil).FindCache), arg0, arg1)
}