package usecases

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"post-service/internal/models"
	repoMock "post-service/internal/repository/mock"
	"testing"
	"time"
)

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repoMock.NewMockRepository(ctrl)

	ctx := context.Background()
	usecase := &commentsUsecase{
		repo: mockRepo,
	}

	args := map[string]interface{}{
		"post_id":   "1",
		"parent_id": "0",
		"content":   "Test Content",
		"author_id": "123",
	}

	mockPost := &models.Post{
		ID:            1,
		AllowComments: true,
	}

	mockComment := &models.Comment{
		PostID:   1,
		ParentID: 0,
		Content:  "Test Content",
		AuthorID: 123,
	}

	mockRepo.EXPECT().GetPost(ctx, int64(1)).Return(mockPost, nil)
	mockRepo.EXPECT().CreateComment(ctx, mockComment).Return(nil)

	comment, err := usecase.CreateComment(ctx, args)

	assert.NoError(t, err)
	assert.Equal(t, mockComment.PostID, comment.PostID)
	assert.Equal(t, mockComment.ParentID, comment.ParentID)
	assert.Equal(t, mockComment.Content, comment.Content)
	assert.Equal(t, mockComment.AuthorID, comment.AuthorID)
}

func TestGetComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repoMock.NewMockRepository(ctrl)

	ctx := context.Background()
	usecase := &commentsUsecase{
		repo: mockRepo,
	}

	args := map[string]interface{}{
		"id": "1",
	}

	mockComment := &models.Comment{
		ID:        1,
		PostID:    1,
		ParentID:  0,
		Content:   "Test Content",
		AuthorID:  123,
		CreatedAt: time.Now(),
	}

	mockRepo.EXPECT().GetComment(ctx, int64(1)).Return(mockComment, nil)

	comment, err := usecase.GetComment(ctx, args)

	assert.NoError(t, err)
	assert.Equal(t, mockComment, comment)
}

func TestGetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repoMock.NewMockRepository(ctrl)

	ctx := context.Background()
	usecase := &commentsUsecase{
		repo: mockRepo,
	}

	mockComments := []*models.Comment{
		{
			ID:        1,
			PostID:    1,
			ParentID:  0,
			Content:   "Test Content 1",
			AuthorID:  123,
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			PostID:    1,
			ParentID:  0,
			Content:   "Test Content 2",
			AuthorID:  456,
			CreatedAt: time.Now(),
		},
	}

	tests := []struct {
		name          string
		args          map[string]interface{}
		mockComments  []*models.Comment
		mockPostFound bool
		expectErr     bool
		mockFunc      func()
	}{
		{
			name: "successful with post_id",
			args: map[string]interface{}{
				"post_id": "1",
				"limit":   10,
				"offset":  0,
			},
			mockComments: []*models.Comment{
				{
					ID:        1,
					PostID:    1,
					ParentID:  0,
					Content:   "Test Content 1",
					AuthorID:  123,
					CreatedAt: time.Now(),
				},
				{
					ID:        2,
					PostID:    1,
					ParentID:  0,
					Content:   "Test Content 2",
					AuthorID:  456,
					CreatedAt: time.Now(),
				},
			},
			mockPostFound: true,
			expectErr:     false,
			mockFunc: func() {
				mockRepo.EXPECT().GetCommentsByPostID(ctx, int64(1), 10, 0).Return(mockComments, nil)
			},
		},
		{
			name: "successful with id",
			args: map[string]interface{}{
				"id":     "1",
				"limit":  10,
				"offset": 0,
			},
			mockComments: []*models.Comment{
				{
					ID:        1,
					PostID:    1,
					ParentID:  0,
					Content:   "Test Content 1",
					AuthorID:  123,
					CreatedAt: time.Now(),
				},
				{
					ID:        2,
					PostID:    1,
					ParentID:  0,
					Content:   "Test Content 2",
					AuthorID:  456,
					CreatedAt: time.Now(),
				},
			},
			mockPostFound: true,
			expectErr:     false,
			mockFunc: func() {
				mockRepo.EXPECT().GetCommentsByPostID(ctx, int64(1), 10, 0).Return(mockComments, nil)
			},
		},
		{
			name: "repository error",
			args: map[string]interface{}{
				"post_id": "1",
				"limit":   10,
				"offset":  0,
			},
			mockComments:  nil,
			mockPostFound: true,
			expectErr:     true,
			mockFunc: func() {
				mockRepo.EXPECT().GetCommentsByPostID(ctx, int64(1), 10, 0).Return(nil, errors.New("repository error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc()
			}

			comments, err := usecase.GetComments(ctx, tt.args)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, comments)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, comments)
				assert.Equal(t, len(tt.mockComments), len(comments))
			}
		})
	}
}
