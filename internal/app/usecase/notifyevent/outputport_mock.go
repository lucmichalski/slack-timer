// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/usecase/notifyevent/outputport.go

// Package notifyevent is a generated GoMock package.
package notifyevent

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOutputPort is a mock of OutputPort interface
type MockOutputPort struct {
	ctrl     *gomock.Controller
	recorder *MockOutputPortMockRecorder
}

// MockOutputPortMockRecorder is the mock recorder for MockOutputPort
type MockOutputPortMockRecorder struct {
	mock *MockOutputPort
}

// NewMockOutputPort creates a new mock instance
func NewMockOutputPort(ctrl *gomock.Controller) *MockOutputPort {
	mock := &MockOutputPort{ctrl: ctrl}
	mock.recorder = &MockOutputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOutputPort) EXPECT() *MockOutputPortMockRecorder {
	return m.recorder
}

// Output mocks base method
func (m *MockOutputPort) Output(data OutputData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Output", data)
}

// Output indicates an expected call of Output
func (mr *MockOutputPortMockRecorder) Output(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockOutputPort)(nil).Output), data)
}