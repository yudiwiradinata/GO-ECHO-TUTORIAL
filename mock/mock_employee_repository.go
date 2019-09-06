// Code generated by MockGen. DO NOT EDIT.
// Source: GO-ECHO-TUTORIAL/repository (interfaces: EmployeeRepositoryInterface)

// Package mock is a generated GoMock package.
package mock

import (
	model "GO-ECHO-TUTORIAL/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmployeeRepositoryInterface is a mock of EmployeeRepositoryInterface interface
type MockEmployeeRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeRepositoryInterfaceMockRecorder
}

// MockEmployeeRepositoryInterfaceMockRecorder is the mock recorder for MockEmployeeRepositoryInterface
type MockEmployeeRepositoryInterfaceMockRecorder struct {
	mock *MockEmployeeRepositoryInterface
}

// NewMockEmployeeRepositoryInterface creates a new mock instance
func NewMockEmployeeRepositoryInterface(ctrl *gomock.Controller) *MockEmployeeRepositoryInterface {
	mock := &MockEmployeeRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockEmployeeRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEmployeeRepositoryInterface) EXPECT() *MockEmployeeRepositoryInterfaceMockRecorder {
	return m.recorder
}

// DeleteEmployee mocks base method
func (m *MockEmployeeRepositoryInterface) DeleteEmployee(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployee", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEmployee indicates an expected call of DeleteEmployee
func (mr *MockEmployeeRepositoryInterfaceMockRecorder) DeleteEmployee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployee", reflect.TypeOf((*MockEmployeeRepositoryInterface)(nil).DeleteEmployee), arg0)
}

// GetEmployee mocks base method
func (m *MockEmployeeRepositoryInterface) GetEmployee() (*model.Employees, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployee")
	ret0, _ := ret[0].(*model.Employees)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployee indicates an expected call of GetEmployee
func (mr *MockEmployeeRepositoryInterfaceMockRecorder) GetEmployee() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployee", reflect.TypeOf((*MockEmployeeRepositoryInterface)(nil).GetEmployee))
}

// GetEmployeeByID mocks base method
func (m *MockEmployeeRepositoryInterface) GetEmployeeByID(arg0 int64) (*model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByID", arg0)
	ret0, _ := ret[0].(*model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByID indicates an expected call of GetEmployeeByID
func (mr *MockEmployeeRepositoryInterfaceMockRecorder) GetEmployeeByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByID", reflect.TypeOf((*MockEmployeeRepositoryInterface)(nil).GetEmployeeByID), arg0)
}

// GetEmployeeByName mocks base method
func (m *MockEmployeeRepositoryInterface) GetEmployeeByName(arg0 string) (*model.Employees, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByName", arg0)
	ret0, _ := ret[0].(*model.Employees)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByName indicates an expected call of GetEmployeeByName
func (mr *MockEmployeeRepositoryInterfaceMockRecorder) GetEmployeeByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByName", reflect.TypeOf((*MockEmployeeRepositoryInterface)(nil).GetEmployeeByName), arg0)
}

// InsertEmployee mocks base method
func (m *MockEmployeeRepositoryInterface) InsertEmployee(arg0 model.EmployeeRequest) (*model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertEmployee", arg0)
	ret0, _ := ret[0].(*model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertEmployee indicates an expected call of InsertEmployee
func (mr *MockEmployeeRepositoryInterfaceMockRecorder) InsertEmployee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEmployee", reflect.TypeOf((*MockEmployeeRepositoryInterface)(nil).InsertEmployee), arg0)
}

// UpdateEmployee mocks base method
func (m *MockEmployeeRepositoryInterface) UpdateEmployee(arg0 string, arg1 model.EmployeeRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEmployee indicates an expected call of UpdateEmployee
func (mr *MockEmployeeRepositoryInterfaceMockRecorder) UpdateEmployee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockEmployeeRepositoryInterface)(nil).UpdateEmployee), arg0, arg1)
}
