package repository

import (
	"context"
	"post-service/internal/models"
)

type Repository interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPost(ctx context.Context, id int64) (*models.Post, error)
	GetPosts(ctx context.Context) ([]*models.Post, error)
	CreateComment(ctx context.Context, comment *models.Comment) error
	GetCommentsByPostID(ctx context.Context, postID int64, limit, offset int) ([]*models.Comment, error)
	GetComment(ctx context.Context, id int64) (*models.Comment, error)
}
