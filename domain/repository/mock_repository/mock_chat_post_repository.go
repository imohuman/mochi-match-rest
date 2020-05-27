// Code generated by MockGen. DO NOT EDIT.
// Source: chat_post_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/taniwhy/mochi-match-rest/domain/models"
	reflect "reflect"
)

// MockChatPostRepository is a mock of ChatPostRepository interface
type MockChatPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatPostRepositoryMockRecorder
}

// MockChatPostRepositoryMockRecorder is the mock recorder for MockChatPostRepository
type MockChatPostRepositoryMockRecorder struct {
	mock *MockChatPostRepository
}

// NewMockChatPostRepository creates a new mock instance
func NewMockChatPostRepository(ctrl *gomock.Controller) *MockChatPostRepository {
	mock := &MockChatPostRepository{ctrl: ctrl}
	mock.recorder = &MockChatPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatPostRepository) EXPECT() *MockChatPostRepositoryMockRecorder {
	return m.recorder
}

// FindAllChatPost mocks base method
func (m *MockChatPostRepository) FindAllChatPost() ([]*models.ChatPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllChatPost")
	ret0, _ := ret[0].([]*models.ChatPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllChatPost indicates an expected call of FindAllChatPost
func (mr *MockChatPostRepositoryMockRecorder) FindAllChatPost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllChatPost", reflect.TypeOf((*MockChatPostRepository)(nil).FindAllChatPost))
}

// FindChatPostByRoomID mocks base method
func (m *MockChatPostRepository) FindChatPostByRoomID(id string) ([]*models.ChatPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindChatPostByRoomID", id)
	ret0, _ := ret[0].([]*models.ChatPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChatPostByRoomID indicates an expected call of FindChatPostByRoomID
func (mr *MockChatPostRepositoryMockRecorder) FindChatPostByRoomID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChatPostByRoomID", reflect.TypeOf((*MockChatPostRepository)(nil).FindChatPostByRoomID), id)
}

// FindChatPostByRoomIDAndLimit mocks base method
func (m *MockChatPostRepository) FindChatPostByRoomIDAndLimit(id string, limit int) ([]*models.ChatPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindChatPostByRoomIDAndLimit", id, limit)
	ret0, _ := ret[0].([]*models.ChatPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChatPostByRoomIDAndLimit indicates an expected call of FindChatPostByRoomIDAndLimit
func (mr *MockChatPostRepositoryMockRecorder) FindChatPostByRoomIDAndLimit(id, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChatPostByRoomIDAndLimit", reflect.TypeOf((*MockChatPostRepository)(nil).FindChatPostByRoomIDAndLimit), id, limit)
}

// FindChatPostByRoomIDAndOffset mocks base method
func (m *MockChatPostRepository) FindChatPostByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindChatPostByRoomIDAndOffset", id, offset)
	ret0, _ := ret[0].([]*models.ChatPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChatPostByRoomIDAndOffset indicates an expected call of FindChatPostByRoomIDAndOffset
func (mr *MockChatPostRepositoryMockRecorder) FindChatPostByRoomIDAndOffset(id, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChatPostByRoomIDAndOffset", reflect.TypeOf((*MockChatPostRepository)(nil).FindChatPostByRoomIDAndOffset), id, offset)
}

// FindChatPostByRoomIDAndLimitAndOffset mocks base method
func (m *MockChatPostRepository) FindChatPostByRoomIDAndLimitAndOffset(id, offset string, limit int) ([]*models.ChatPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindChatPostByRoomIDAndLimitAndOffset", id, offset, limit)
	ret0, _ := ret[0].([]*models.ChatPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChatPostByRoomIDAndLimitAndOffset indicates an expected call of FindChatPostByRoomIDAndLimitAndOffset
func (mr *MockChatPostRepositoryMockRecorder) FindChatPostByRoomIDAndLimitAndOffset(id, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChatPostByRoomIDAndLimitAndOffset", reflect.TypeOf((*MockChatPostRepository)(nil).FindChatPostByRoomIDAndLimitAndOffset), id, offset, limit)
}

// InsertChatPost mocks base method
func (m *MockChatPostRepository) InsertChatPost(room *models.ChatPost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertChatPost", room)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertChatPost indicates an expected call of InsertChatPost
func (mr *MockChatPostRepositoryMockRecorder) InsertChatPost(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertChatPost", reflect.TypeOf((*MockChatPostRepository)(nil).InsertChatPost), room)
}

// DeleteChatPost mocks base method
func (m *MockChatPostRepository) DeleteChatPost(room *models.ChatPost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChatPost", room)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChatPost indicates an expected call of DeleteChatPost
func (mr *MockChatPostRepositoryMockRecorder) DeleteChatPost(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChatPost", reflect.TypeOf((*MockChatPostRepository)(nil).DeleteChatPost), room)
}