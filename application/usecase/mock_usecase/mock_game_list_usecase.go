// Code generated by MockGen. DO NOT EDIT.
// Source: game_list_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	models "github.com/taniwhy/mochi-match-rest/domain/models"
	reflect "reflect"
)

// MockIGameListUseCase is a mock of IGameListUseCase interface
type MockIGameListUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIGameListUseCaseMockRecorder
}

// MockIGameListUseCaseMockRecorder is the mock recorder for MockIGameListUseCase
type MockIGameListUseCaseMockRecorder struct {
	mock *MockIGameListUseCase
}

// NewMockIGameListUseCase creates a new mock instance
func NewMockIGameListUseCase(ctrl *gomock.Controller) *MockIGameListUseCase {
	mock := &MockIGameListUseCase{ctrl: ctrl}
	mock.recorder = &MockIGameListUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGameListUseCase) EXPECT() *MockIGameListUseCaseMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockIGameListUseCase) FindAll(c *gin.Context) ([]*models.GameList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", c)
	ret0, _ := ret[0].([]*models.GameList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockIGameListUseCaseMockRecorder) FindAll(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIGameListUseCase)(nil).FindAll), c)
}

// Insert mocks base method
func (m *MockIGameListUseCase) Insert(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIGameListUseCaseMockRecorder) Insert(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIGameListUseCase)(nil).Insert), c)
}

// Update mocks base method
func (m *MockIGameListUseCase) Update(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockIGameListUseCaseMockRecorder) Update(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIGameListUseCase)(nil).Update), c)
}

// Delete mocks base method
func (m *MockIGameListUseCase) Delete(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIGameListUseCaseMockRecorder) Delete(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIGameListUseCase)(nil).Delete), c)
}