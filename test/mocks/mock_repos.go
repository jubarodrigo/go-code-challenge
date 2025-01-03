// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	actions "go-code-challenge/internal/actions"
	users "go-code-challenge/internal/users"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDatasJsonRepositoryInterface is a mock of DatasJsonRepositoryInterface interface.
type MockDatasJsonRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDatasJsonRepositoryInterfaceMockRecorder
}

// MockDatasJsonRepositoryInterfaceMockRecorder is the mock recorder for MockDatasJsonRepositoryInterface.
type MockDatasJsonRepositoryInterfaceMockRecorder struct {
	mock *MockDatasJsonRepositoryInterface
}

// NewMockDatasJsonRepositoryInterface creates a new mock instance.
func NewMockDatasJsonRepositoryInterface(ctrl *gomock.Controller) *MockDatasJsonRepositoryInterface {
	mock := &MockDatasJsonRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockDatasJsonRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatasJsonRepositoryInterface) EXPECT() *MockDatasJsonRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetActionsByUserID mocks base method.
func (m *MockDatasJsonRepositoryInterface) GetActionsByUserID(userID int) ([]actions.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionsByUserID", userID)
	ret0, _ := ret[0].([]actions.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionsByUserID indicates an expected call of GetActionsByUserID.
func (mr *MockDatasJsonRepositoryInterfaceMockRecorder) GetActionsByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionsByUserID", reflect.TypeOf((*MockDatasJsonRepositoryInterface)(nil).GetActionsByUserID), userID)
}

// GetAllActions mocks base method.
func (m *MockDatasJsonRepositoryInterface) GetAllActions() ([]actions.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllActions")
	ret0, _ := ret[0].([]actions.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllActions indicates an expected call of GetAllActions.
func (mr *MockDatasJsonRepositoryInterfaceMockRecorder) GetAllActions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllActions", reflect.TypeOf((*MockDatasJsonRepositoryInterface)(nil).GetAllActions))
}

// GetUserByID mocks base method.
func (m *MockDatasJsonRepositoryInterface) GetUserByID(id int) (*users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockDatasJsonRepositoryInterfaceMockRecorder) GetUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDatasJsonRepositoryInterface)(nil).GetUserByID), id)
}
