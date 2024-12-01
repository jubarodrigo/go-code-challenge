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

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserByID mocks base method.
func (m *MockUserRepository) GetUserByID(id int) (*users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryMockRecorder) GetUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserByID), id)
}

// MockActionRepository is a mock of ActionRepository interface.
type MockActionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockActionRepositoryMockRecorder
}

// MockActionRepositoryMockRecorder is the mock recorder for MockActionRepository.
type MockActionRepositoryMockRecorder struct {
	mock *MockActionRepository
}

// NewMockActionRepository creates a new mock instance.
func NewMockActionRepository(ctrl *gomock.Controller) *MockActionRepository {
	mock := &MockActionRepository{ctrl: ctrl}
	mock.recorder = &MockActionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionRepository) EXPECT() *MockActionRepositoryMockRecorder {
	return m.recorder
}

// GetActionsByUserID mocks base method.
func (m *MockActionRepository) GetActionsByUserID(userID int) ([]actions.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionsByUserID", userID)
	ret0, _ := ret[0].([]actions.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionsByUserID indicates an expected call of GetActionsByUserID.
func (mr *MockActionRepositoryMockRecorder) GetActionsByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionsByUserID", reflect.TypeOf((*MockActionRepository)(nil).GetActionsByUserID), userID)
}

// GetAllActions mocks base method.
func (m *MockActionRepository) GetAllActions() ([]actions.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllActions")
	ret0, _ := ret[0].([]actions.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllActions indicates an expected call of GetAllActions.
func (mr *MockActionRepositoryMockRecorder) GetAllActions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllActions", reflect.TypeOf((*MockActionRepository)(nil).GetAllActions))
}