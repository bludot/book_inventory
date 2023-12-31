// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bludot/tempmee/inventory/internal/db/repositories/book (interfaces: BookRepositoryImpl)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_book_repository.go -package=mocks github.com/bludot/tempmee/inventory/internal/db/repositories/book BookRepositoryImpl
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	book "github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	gomock "go.uber.org/mock/gomock"
)

// MockBookRepositoryImpl is a mock of BookRepositoryImpl interface.
type MockBookRepositoryImpl struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryImplMockRecorder
}

// MockBookRepositoryImplMockRecorder is the mock recorder for MockBookRepositoryImpl.
type MockBookRepositoryImplMockRecorder struct {
	mock *MockBookRepositoryImpl
}

// NewMockBookRepositoryImpl creates a new mock instance.
func NewMockBookRepositoryImpl(ctrl *gomock.Controller) *MockBookRepositoryImpl {
	mock := &MockBookRepositoryImpl{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepositoryImpl) EXPECT() *MockBookRepositoryImplMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBookRepositoryImpl) Create(arg0 context.Context, arg1 book.Book) (*book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBookRepositoryImplMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookRepositoryImpl)(nil).Create), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockBookRepositoryImpl) FindAll(arg0 context.Context) (*[]book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].(*[]book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockBookRepositoryImplMockRecorder) FindAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockBookRepositoryImpl)(nil).FindAll), arg0)
}

// FindByAuthor mocks base method.
func (m *MockBookRepositoryImpl) FindByAuthor(arg0 context.Context, arg1 string) (*[]book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAuthor", arg0, arg1)
	ret0, _ := ret[0].(*[]book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAuthor indicates an expected call of FindByAuthor.
func (mr *MockBookRepositoryImplMockRecorder) FindByAuthor(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAuthor", reflect.TypeOf((*MockBookRepositoryImpl)(nil).FindByAuthor), arg0, arg1)
}

// FindById mocks base method.
func (m *MockBookRepositoryImpl) FindById(arg0 context.Context, arg1 string) (*book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockBookRepositoryImplMockRecorder) FindById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockBookRepositoryImpl)(nil).FindById), arg0, arg1)
}

// FindByTitle mocks base method.
func (m *MockBookRepositoryImpl) FindByTitle(arg0 context.Context, arg1 string) (*book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTitle", arg0, arg1)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTitle indicates an expected call of FindByTitle.
func (mr *MockBookRepositoryImplMockRecorder) FindByTitle(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTitle", reflect.TypeOf((*MockBookRepositoryImpl)(nil).FindByTitle), arg0, arg1)
}

// GetBooksByIDs mocks base method.
func (m *MockBookRepositoryImpl) GetBooksByIDs(arg0 context.Context, arg1 []string) (*[]book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByIDs", arg0, arg1)
	ret0, _ := ret[0].(*[]book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooksByIDs indicates an expected call of GetBooksByIDs.
func (mr *MockBookRepositoryImplMockRecorder) GetBooksByIDs(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByIDs", reflect.TypeOf((*MockBookRepositoryImpl)(nil).GetBooksByIDs), arg0, arg1)
}
