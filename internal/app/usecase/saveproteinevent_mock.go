// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/usecase/saveproteinevent.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProteinEventSaver is a mock of ProteinEventSaver interface
type MockProteinEventSaver struct {
	ctrl     *gomock.Controller
	recorder *MockProteinEventSaverMockRecorder
}

// MockProteinEventSaverMockRecorder is the mock recorder for MockProteinEventSaver
type MockProteinEventSaverMockRecorder struct {
	mock *MockProteinEventSaver
}

// NewMockProteinEventSaver creates a new mock instance
func NewMockProteinEventSaver(ctrl *gomock.Controller) *MockProteinEventSaver {
	mock := &MockProteinEventSaver{ctrl: ctrl}
	mock.recorder = &MockProteinEventSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProteinEventSaver) EXPECT() *MockProteinEventSaverMockRecorder {
	return m.recorder
}

// UpdateTimeToDrink mocks base method
func (m *MockProteinEventSaver) UpdateTimeToDrink(ctx context.Context, userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTimeToDrink", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTimeToDrink indicates an expected call of UpdateTimeToDrink
func (mr *MockProteinEventSaverMockRecorder) UpdateTimeToDrink(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimeToDrink", reflect.TypeOf((*MockProteinEventSaver)(nil).UpdateTimeToDrink), ctx, userId)
}

// SaveIntervalMin mocks base method
func (m *MockProteinEventSaver) SaveIntervalMin(ctx context.Context, userId string, minutes int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveIntervalMin", ctx, userId, minutes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveIntervalMin indicates an expected call of SaveIntervalMin
func (mr *MockProteinEventSaverMockRecorder) SaveIntervalMin(ctx, userId, minutes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveIntervalMin", reflect.TypeOf((*MockProteinEventSaver)(nil).SaveIntervalMin), ctx, userId, minutes)
}
