// Code generated by MockGen. DO NOT EDIT.
// Source: report_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIReportUsecase is a mock of IReportUsecase interface
type MockIReportUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIReportUsecaseMockRecorder
}

// MockIReportUsecaseMockRecorder is the mock recorder for MockIReportUsecase
type MockIReportUsecaseMockRecorder struct {
	mock *MockIReportUsecase
}

// NewMockIReportUsecase creates a new mock instance
func NewMockIReportUsecase(ctrl *gomock.Controller) *MockIReportUsecase {
	mock := &MockIReportUsecase{ctrl: ctrl}
	mock.recorder = &MockIReportUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIReportUsecase) EXPECT() *MockIReportUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIReportUsecase) Create(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockIReportUsecaseMockRecorder) Create(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIReportUsecase)(nil).Create), c)
}
