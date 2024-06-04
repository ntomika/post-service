package usecases

import (
	"context"
	"errors"
	"post-service/internal/models"
	"post-service/internal/repository"
	"strconv"
)

type postUsecase struct {
	repo repository.Repository
}

func NewPostsUsecase(repo repository.Repository) Posts {
	return &postUsecase{repo: repo}
}

// CreatePost creates a new post
func (u *postUsecase) CreatePost(ctx context.Context, args map[string]interface{}) (*models.Post, error) {
	title, ok := args["title"].(string)
	if !ok {
		return nil, errors.New("missing or invalid title")
	}

	content, ok := args["content"].(string)
	if !ok {
		return nil, errors.New("missing or invalid content")
	}

	authorIDStr, ok := args["author_id"].(string)
	if !ok {
		return nil, errors.New("missing or invalid author_id")
	}

	authorID, err := strconv.ParseInt(authorIDStr, 10, 64)
	if err != nil {
		return nil, err
	}

	allowComments, ok := args["allow_comments"].(bool)
	if !ok {
		return nil, errors.New("missing or invalid allow_comments")
	}

	post := &models.Post{
		Title:         title,
		Content:       content,
		AuthorID:      authorID,
		AllowComments: allowComments,
	}

	err = u.repo.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetPost get post by ID
func (u *postUsecase) GetPost(ctx context.Context, args map[string]interface{}) (*models.Post, error) {
	id, err := strconv.ParseInt(args["id"].(string), 10, 64)
	if err != nil {
		return nil, err
	}

	post, err := u.repo.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetPosts gets list of posts
func (u *postUsecase) GetPosts(ctx context.Context) ([]*models.Post, error) {
	posts, err := u.repo.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
