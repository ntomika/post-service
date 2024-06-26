// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/interface.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/interface.go -destination=internal/repository/mock/repository.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "post-service/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockRepository) CreateComment(ctx context.Context, comment *models.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockRepositoryMockRecorder) CreateComment(ctx, comment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockRepository)(nil).CreateComment), ctx, comment)
}

// CreatePost mocks base method.
func (m *MockRepository) CreatePost(ctx context.Context, post *models.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockRepositoryMockRecorder) CreatePost(ctx, post any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockRepository)(nil).CreatePost), ctx, post)
}

// GetComment mocks base method.
func (m *MockRepository) GetComment(ctx context.Context, id int64) (*models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", ctx, id)
	ret0, _ := ret[0].(*models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockRepositoryMockRecorder) GetComment(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockRepository)(nil).GetComment), ctx, id)
}

// GetCommentsByPostID mocks base method.
func (m *MockRepository) GetCommentsByPostID(ctx context.Context, postID int64, limit, offset int) ([]*models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsByPostID", ctx, postID, limit, offset)
	ret0, _ := ret[0].([]*models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsByPostID indicates an expected call of GetCommentsByPostID.
func (mr *MockRepositoryMockRecorder) GetCommentsByPostID(ctx, postID, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsByPostID", reflect.TypeOf((*MockRepository)(nil).GetCommentsByPostID), ctx, postID, limit, offset)
}

// GetPost mocks base method.
func (m *MockRepository) GetPost(ctx context.Context, id int64) (*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", ctx, id)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockRepositoryMockRecorder) GetPost(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockRepository)(nil).GetPost), ctx, id)
}

// GetPosts mocks base method.
func (m *MockRepository) GetPosts(ctx context.Context) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", ctx)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockRepositoryMockRecorder) GetPosts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockRepository)(nil).GetPosts), ctx)
}
