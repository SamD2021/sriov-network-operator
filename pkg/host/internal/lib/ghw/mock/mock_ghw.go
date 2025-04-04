// Code generated by MockGen. DO NOT EDIT.
// Source: ghw.go
//
// Generated by this command:
//
//	mockgen -destination mock/mock_ghw.go -source ghw.go
//

// Package mock_ghw is a generated GoMock package.
package mock_ghw

import (
	reflect "reflect"

	cpu "github.com/jaypipes/ghw/pkg/cpu"
	pci "github.com/jaypipes/ghw/pkg/pci"
	gomock "go.uber.org/mock/gomock"
)

// MockGHWLib is a mock of GHWLib interface.
type MockGHWLib struct {
	ctrl     *gomock.Controller
	recorder *MockGHWLibMockRecorder
	isgomock struct{}
}

// MockGHWLibMockRecorder is the mock recorder for MockGHWLib.
type MockGHWLibMockRecorder struct {
	mock *MockGHWLib
}

// NewMockGHWLib creates a new mock instance.
func NewMockGHWLib(ctrl *gomock.Controller) *MockGHWLib {
	mock := &MockGHWLib{ctrl: ctrl}
	mock.recorder = &MockGHWLibMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGHWLib) EXPECT() *MockGHWLibMockRecorder {
	return m.recorder
}

// CPU mocks base method.
func (m *MockGHWLib) CPU() (*cpu.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CPU")
	ret0, _ := ret[0].(*cpu.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CPU indicates an expected call of CPU.
func (mr *MockGHWLibMockRecorder) CPU() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CPU", reflect.TypeOf((*MockGHWLib)(nil).CPU))
}

// PCI mocks base method.
func (m *MockGHWLib) PCI() (*pci.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PCI")
	ret0, _ := ret[0].(*pci.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PCI indicates an expected call of PCI.
func (mr *MockGHWLibMockRecorder) PCI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PCI", reflect.TypeOf((*MockGHWLib)(nil).PCI))
}
