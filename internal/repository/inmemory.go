package repository

import (
	"context"
	"errors"
	"post-service/internal/models"
	"time"
)

func (repo *inMemoryRepo) CreatePost(_ context.Context, post *models.Post) error {
	post.CreatedAt = time.Now()

	repo.mu.Lock()
	defer repo.mu.Unlock()

	post.ID = int64(len(repo.posts) + 1)
	repo.posts = append(repo.posts, post)

	return nil
}

func (repo *inMemoryRepo) GetPosts(_ context.Context) ([]*models.Post, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	return repo.posts, nil
}

func (repo *inMemoryRepo) GetPost(_ context.Context, id int64) (*models.Post, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, post := range repo.posts {
		if post.ID == id {
			return post, nil
		}
	}

	return nil, errors.New("post not found")
}

func (repo *inMemoryRepo) CreateComment(_ context.Context, comment *models.Comment) error {
	comment.CreatedAt = time.Now()

	repo.mu.Lock()
	defer repo.mu.Unlock()

	comment.ID = int64(len(repo.comments) + 1)
	repo.comments = append(repo.comments, comment)

	return nil
}

func (repo *inMemoryRepo) GetCommentsByPostID(
	_ context.Context,
	postID int64,
	limit, offset int,
) ([]*models.Comment, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var comments []*models.Comment
	for _, comment := range repo.comments {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}

	if offset > len(comments) {
		return nil, errors.New("offset exceeds the number of comments")
	}

	end := offset + limit
	if end > len(comments) {
		end = len(comments)
	}

	return comments[offset:end], nil
}

func (repo *inMemoryRepo) GetComment(_ context.Context, id int64) (*models.Comment, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, comment := range repo.comments {
		if comment.ID == id {
			return comment, nil
		}
	}
	return nil, errors.New("comment not found")
}
