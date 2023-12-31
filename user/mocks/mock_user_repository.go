// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bludot/tempmee/user/internal/db/repositories/user (interfaces: UserRepositoryImpl)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_user_repository.go -package=mocks github.com/bludot/tempmee/user/internal/db/repositories/user UserRepositoryImpl
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	user "github.com/bludot/tempmee/user/internal/db/repositories/user"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepositoryImpl is a mock of UserRepositoryImpl interface.
type MockUserRepositoryImpl struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryImplMockRecorder
}

// MockUserRepositoryImplMockRecorder is the mock recorder for MockUserRepositoryImpl.
type MockUserRepositoryImplMockRecorder struct {
	mock *MockUserRepositoryImpl
}

// NewMockUserRepositoryImpl creates a new mock instance.
func NewMockUserRepositoryImpl(ctrl *gomock.Controller) *MockUserRepositoryImpl {
	mock := &MockUserRepositoryImpl{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryImpl) EXPECT() *MockUserRepositoryImplMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepositoryImpl) CreateUser(arg0, arg1, arg2 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryImplMockRecorder) CreateUser(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepositoryImpl)(nil).CreateUser), arg0, arg1, arg2)
}

// FindAll mocks base method.
func (m *MockUserRepositoryImpl) FindAll() ([]*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserRepositoryImplMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserRepositoryImpl)(nil).FindAll))
}

// FindByEmail mocks base method.
func (m *MockUserRepositoryImpl) FindByEmail(arg0 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryImplMockRecorder) FindByEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepositoryImpl)(nil).FindByEmail), arg0)
}

// FindById mocks base method.
func (m *MockUserRepositoryImpl) FindById(arg0 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserRepositoryImplMockRecorder) FindById(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserRepositoryImpl)(nil).FindById), arg0)
}

// FindByName mocks base method.
func (m *MockUserRepositoryImpl) FindByName(arg0 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockUserRepositoryImplMockRecorder) FindByName(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockUserRepositoryImpl)(nil).FindByName), arg0)
}
