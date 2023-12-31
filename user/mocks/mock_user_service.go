// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bludot/tempmee/user/internal/services/user (interfaces: UserServiceImpl)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_user_service.go -package=mocks github.com/bludot/tempmee/user/internal/services/user UserServiceImpl
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	user "github.com/bludot/tempmee/user/internal/db/repositories/user"
	gomock "go.uber.org/mock/gomock"
)

// MockUserServiceImpl is a mock of UserServiceImpl interface.
type MockUserServiceImpl struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceImplMockRecorder
}

// MockUserServiceImplMockRecorder is the mock recorder for MockUserServiceImpl.
type MockUserServiceImplMockRecorder struct {
	mock *MockUserServiceImpl
}

// NewMockUserServiceImpl creates a new mock instance.
func NewMockUserServiceImpl(ctrl *gomock.Controller) *MockUserServiceImpl {
	mock := &MockUserServiceImpl{ctrl: ctrl}
	mock.recorder = &MockUserServiceImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceImpl) EXPECT() *MockUserServiceImplMockRecorder {
	return m.recorder
}

// GetAllUsers mocks base method.
func (m *MockUserServiceImpl) GetAllUsers(arg0 context.Context, arg1 int) ([]*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0, arg1)
	ret0, _ := ret[0].([]*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserServiceImplMockRecorder) GetAllUsers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserServiceImpl)(nil).GetAllUsers), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockUserServiceImpl) GetUserByEmail(arg0 context.Context, arg1 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserServiceImplMockRecorder) GetUserByEmail(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserServiceImpl)(nil).GetUserByEmail), arg0, arg1)
}

// Register mocks base method.
func (m *MockUserServiceImpl) Register(arg0 context.Context, arg1, arg2, arg3 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceImplMockRecorder) Register(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserServiceImpl)(nil).Register), arg0, arg1, arg2, arg3)
}

// SignIn mocks base method.
func (m *MockUserServiceImpl) SignIn(arg0 context.Context, arg1, arg2 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", arg0, arg1, arg2)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserServiceImplMockRecorder) SignIn(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUserServiceImpl)(nil).SignIn), arg0, arg1, arg2)
}
