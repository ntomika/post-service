package usecases

import (
	"context"
	"post-service/internal/models"
)

type Posts interface {
	CreatePost(ctx context.Context, args map[string]interface{}) (*models.Post, error)
	GetPost(ctx context.Context, args map[string]interface{}) (*models.Post, error)
	GetPosts(ctx context.Context) ([]*models.Post, error)
}

type Comments interface {
	CreateComment(ctx context.Context, args map[string]interface{}) (*models.Comment, error)
	GetComment(ctx context.Context, args map[string]interface{}) (*models.Comment, error)
	GetComments(ctx context.Context, args map[string]interface{}) ([]*models.Comment, error)
}
