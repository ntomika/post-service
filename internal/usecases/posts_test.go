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

func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repoMock.NewMockRepository(ctrl)

	ctx := context.Background()
	usecase := &postUsecase{
		repo: mockRepo,
	}

	tests := []struct {
		name      string
		args      map[string]interface{}
		mockPost  *models.Post
		expectErr bool
	}{
		{
			name: "successful creation",
			args: map[string]interface{}{
				"title":          "Test Title",
				"content":        "Test Content",
				"author_id":      "123",
				"allow_comments": true,
			},
			mockPost: &models.Post{
				ID:            0,
				Title:         "Test Title",
				Content:       "Test Content",
				AuthorID:      123,
				AllowComments: true,
			},
			expectErr: false,
		},
		{
			name: "error in parsing author_id",
			args: map[string]interface{}{
				"title":          "Test Title",
				"content":        "Test Content",
				"author_id":      "invalid",
				"allow_comments": true,
			},
			mockPost:  nil,
			expectErr: true,
		},
		{
			name: "missing title",
			args: map[string]interface{}{
				"content":        "Test Content",
				"author_id":      "123",
				"allow_comments": true,
			},
			mockPost:  nil,
			expectErr: true,
		},
		{
			name: "missing content",
			args: map[string]interface{}{
				"title":          "Test Title",
				"author_id":      "123",
				"allow_comments": true,
			},
			mockPost:  nil,
			expectErr: true,
		},
		{
			name: "missing author_id",
			args: map[string]interface{}{
				"title":          "Test Title",
				"content":        "Test Content",
				"allow_comments": true,
			},
			mockPost:  nil,
			expectErr: true,
		},
		{
			name: "repository create post error",
			args: map[string]interface{}{
				"title":          "Test Title",
				"content":        "Test Content",
				"author_id":      "123",
				"allow_comments": true,
			},
			mockPost: &models.Post{
				ID:            0,
				Title:         "Test Title",
				Content:       "Test Content",
				AuthorID:      123,
				AllowComments: true,
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockPost != nil && !tt.expectErr {
				mockRepo.EXPECT().CreatePost(ctx, tt.mockPost).Return(nil)
			} else if tt.mockPost != nil && tt.expectErr {
				mockRepo.EXPECT().CreatePost(ctx, tt.mockPost).Return(errors.New("repository error"))
			}

			post, err := usecase.CreatePost(ctx, tt.args)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, post)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, post)
				assert.Equal(t, tt.mockPost.Title, post.Title)
				assert.Equal(t, tt.mockPost.Content, post.Content)
				assert.Equal(t, tt.mockPost.AuthorID, post.AuthorID)
				assert.Equal(t, tt.mockPost.AllowComments, post.AllowComments)
				assert.NotNil(t, post.CreatedAt)
				assert.Equal(t, int64(0), post.ID)
			}
		})
	}
}

func TestGetPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repoMock.NewMockRepository(ctrl)

	ctx := context.Background()
	usecase := &postUsecase{
		repo: mockRepo,
	}

	tests := []struct {
		name      string
		args      map[string]interface{}
		mockPost  *models.Post
		expectErr bool
	}{
		{
			name: "successful get post",
			args: map[string]interface{}{
				"id": "1",
			},
			mockPost: &models.Post{
				ID:            1,
				Title:         "Test Title",
				Content:       "Test Content",
				AuthorID:      123,
				AllowComments: true,
			},
			expectErr: false,
		},
		{
			name: "repository get post error",
			args: map[string]interface{}{
				"id": "1",
			},
			mockPost:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockPost != nil && !tt.expectErr {
				mockRepo.EXPECT().GetPost(ctx, int64(1)).Return(tt.mockPost, nil)
			} else if tt.mockPost == nil && tt.expectErr {
				mockRepo.EXPECT().GetPost(ctx, int64(1)).Return(nil, errors.New("repository error"))
			}

			post, err := usecase.GetPost(ctx, tt.args)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, post)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockPost, post)
			}
		})
	}
}

func TestGetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repoMock.NewMockRepository(ctrl)

	ctx := context.Background()
	usecase := &postUsecase{
		repo: mockRepo,
	}

	tests := []struct {
		name      string
		mockPosts []*models.Post
		expectErr bool
	}{
		{
			name: "successful get posts",
			mockPosts: []*models.Post{
				{
					ID:            1,
					Title:         "Test Title 1",
					Content:       "Test Content 1",
					AuthorID:      123,
					AllowComments: true,
					CreatedAt:     time.Now(),
				},
				{
					ID:            2,
					Title:         "Test Title 2",
					Content:       "Test Content 2",
					AuthorID:      456,
					AllowComments: false,
					CreatedAt:     time.Now(),
				},
			},
			expectErr: false,
		},
		{
			name:      "repository get posts error",
			mockPosts: nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectErr {
				mockRepo.EXPECT().GetPosts(ctx).Return(tt.mockPosts, nil)
			} else {
				mockRepo.EXPECT().GetPosts(ctx).Return(nil, errors.New("repository error"))
			}

			posts, err := usecase.GetPosts(ctx)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, posts)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockPosts, posts)
			}
		})
	}
}
