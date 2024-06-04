package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"post-service/internal/models"
)

type postgresRepo struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(ctx context.Context, dsn string) (Repository, error) {
	if dsn == "" {
		return nil, errors.New("")
	}

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &postgresRepo{db: conn}, nil
}

type inMemoryRepo struct {
	posts    []*models.Post
	comments []*models.Comment
	mu       sync.RWMutex
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepo{
		posts:    []*models.Post{},
		comments: []*models.Comment{},
	}
}
