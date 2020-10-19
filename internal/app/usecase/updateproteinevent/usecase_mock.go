// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/usecase/updateproteinevent/usecase.go

// Package updateproteinevent is a generated GoMock package.
package updateproteinevent

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// UpdateTimeToDrink mocks base method
func (m *MockUsecase) UpdateTimeToDrink(ctx context.Context, userId string, overWriteOutputPort OutputPort) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateTimeToDrink", ctx, userId, overWriteOutputPort)
}

// UpdateTimeToDrink indicates an expected call of UpdateTimeToDrink
func (mr *MockUsecaseMockRecorder) UpdateTimeToDrink(ctx, userId, overWriteOutputPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimeToDrink", reflect.TypeOf((*MockUsecase)(nil).UpdateTimeToDrink), ctx, userId, overWriteOutputPort)
}

// SaveIntervalMin mocks base method
func (m *MockUsecase) SaveIntervalMin(ctx context.Context, userId string, minutes int, overWriteOutputPort OutputPort) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SaveIntervalMin", ctx, userId, minutes, overWriteOutputPort)
}

// SaveIntervalMin indicates an expected call of SaveIntervalMin
func (mr *MockUsecaseMockRecorder) SaveIntervalMin(ctx, userId, minutes, overWriteOutputPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveIntervalMin", reflect.TypeOf((*MockUsecase)(nil).SaveIntervalMin), ctx, userId, minutes, overWriteOutputPort)
}

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
func (m *MockOutputPort) Output(data *OutputData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Output", data)
}

// Output indicates an expected call of Output
func (mr *MockOutputPortMockRecorder) Output(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockOutputPort)(nil).Output), data)
}
