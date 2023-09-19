// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bludot/tempmee/order/internal/db/repositories/order (interfaces: OrderRepositoryImpl)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_order_repository.go -package=mocks github.com/bludot/tempmee/order/internal/db/repositories/order OrderRepositoryImpl
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	order "github.com/bludot/tempmee/order/internal/db/repositories/order"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderRepositoryImpl is a mock of OrderRepositoryImpl interface.
type MockOrderRepositoryImpl struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryImplMockRecorder
}

// MockOrderRepositoryImplMockRecorder is the mock recorder for MockOrderRepositoryImpl.
type MockOrderRepositoryImplMockRecorder struct {
	mock *MockOrderRepositoryImpl
}

// NewMockOrderRepositoryImpl creates a new mock instance.
func NewMockOrderRepositoryImpl(ctrl *gomock.Controller) *MockOrderRepositoryImpl {
	mock := &MockOrderRepositoryImpl{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepositoryImpl) EXPECT() *MockOrderRepositoryImplMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderRepositoryImpl) Create(arg0 context.Context, arg1 *order.Order) (*order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderRepositoryImplMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockOrderRepositoryImpl) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrderRepositoryImplMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).Delete), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockOrderRepositoryImpl) FindAll(arg0 context.Context) (*[]order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].(*[]order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockOrderRepositoryImplMockRecorder) FindAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).FindAll), arg0)
}

// FindAllByParentOrderId mocks base method.
func (m *MockOrderRepositoryImpl) FindAllByParentOrderId(arg0 context.Context, arg1 int) (*[]order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllByParentOrderId", arg0, arg1)
	ret0, _ := ret[0].(*[]order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByParentOrderId indicates an expected call of FindAllByParentOrderId.
func (mr *MockOrderRepositoryImplMockRecorder) FindAllByParentOrderId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByParentOrderId", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).FindAllByParentOrderId), arg0, arg1)
}

// FindAllByUserId mocks base method.
func (m *MockOrderRepositoryImpl) FindAllByUserId(arg0 context.Context, arg1 string) (*[]order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllByUserId", arg0, arg1)
	ret0, _ := ret[0].(*[]order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByUserId indicates an expected call of FindAllByUserId.
func (mr *MockOrderRepositoryImplMockRecorder) FindAllByUserId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByUserId", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).FindAllByUserId), arg0, arg1)
}

// FindById mocks base method.
func (m *MockOrderRepositoryImpl) FindById(arg0 context.Context, arg1 int) (*order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockOrderRepositoryImplMockRecorder) FindById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).FindById), arg0, arg1)
}

// Update mocks base method.
func (m *MockOrderRepositoryImpl) Update(arg0 context.Context, arg1 *order.Order) (*order.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*order.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrderRepositoryImplMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).Update), arg0, arg1)
}