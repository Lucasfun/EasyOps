// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Lucasfun/EasyOps/unitsInstall_2.0/interface (interfaces: UnitInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUnitInterface is a mock of UnitInterface interface
type MockUnitInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUnitInterfaceMockRecorder
}

// MockUnitInterfaceMockRecorder is the mock recorder for MockUnitInterface
type MockUnitInterfaceMockRecorder struct {
	mock *MockUnitInterface
}

// NewMockUnitInterface creates a new mock instance
func NewMockUnitInterface(ctrl *gomock.Controller) *MockUnitInterface {
	mock := &MockUnitInterface{ctrl: ctrl}
	mock.recorder = &MockUnitInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnitInterface) EXPECT() *MockUnitInterfaceMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockUnitInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockUnitInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockUnitInterface)(nil).GetName))
}

// GetNext mocks base method
func (m *MockUnitInterface) GetNext() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNext")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetNext indicates an expected call of GetNext
func (mr *MockUnitInterfaceMockRecorder) GetNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockUnitInterface)(nil).GetNext))
}

// InstallFunc mocks base method
func (m *MockUnitInterface) InstallFunc(arg0 func()) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallFunc", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// InstallFunc indicates an expected call of InstallFunc
func (mr *MockUnitInterfaceMockRecorder) InstallFunc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallFunc", reflect.TypeOf((*MockUnitInterface)(nil).InstallFunc), arg0)
}
