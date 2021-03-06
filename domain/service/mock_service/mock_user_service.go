// Code generated by MockGen. DO NOT EDIT.
// Source: user_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUserService is a mock of IUserService interface
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// IsAdmin mocks base method
func (m *MockIUserService) IsAdmin(id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAdmin", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAdmin indicates an expected call of IsAdmin
func (mr *MockIUserServiceMockRecorder) IsAdmin(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAdmin", reflect.TypeOf((*MockIUserService)(nil).IsAdmin), id)
}

// IsDelete mocks base method
func (m *MockIUserService) IsDelete(id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDelete", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDelete indicates an expected call of IsDelete
func (mr *MockIUserServiceMockRecorder) IsDelete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDelete", reflect.TypeOf((*MockIUserService)(nil).IsDelete), id)
}

// IsExist mocks base method
func (m *MockIUserService) IsExist(id, provider string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist", id, provider)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExist indicates an expected call of IsExist
func (mr *MockIUserServiceMockRecorder) IsExist(id, provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockIUserService)(nil).IsExist), id, provider)
}
