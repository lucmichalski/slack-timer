// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/driver/di/di.go

// Package di is a generated GoMock package.
package di

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDI is a mock of DI interface
type MockDI struct {
	ctrl     *gomock.Controller
	recorder *MockDIMockRecorder
}

// MockDIMockRecorder is the mock recorder for MockDI
type MockDIMockRecorder struct {
	mock *MockDI
}

// NewMockDI creates a new mock instance
func NewMockDI(ctrl *gomock.Controller) *MockDI {
	mock := &MockDI{ctrl: ctrl}
	mock.recorder = &MockDIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDI) EXPECT() *MockDIMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockDI) Get(name string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockDIMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDI)(nil).Get), name)
}
