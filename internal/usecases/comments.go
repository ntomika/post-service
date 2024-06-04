package usecases

import (
	"context"
	"errors"
	"log"
	"post-service/internal/models"
	"post-service/internal/repository"
	"strconv"
)

const (
	defaultOffset = 0
	defaultLimit  = 10
)

type commentsUsecase struct {
	repo repository.Repository
}

func NewComentsUsecase(repo repository.Repository) Comments {
	return &commentsUsecase{repo: repo}
}

// CreateComment create comment
func (u *commentsUsecase) CreateComment(ctx context.Context, args map[string]interface{}) (*models.Comment, error) {
	postID, err := strconv.ParseInt(args["post_id"].(string), 10, 64)
	if err != nil {
		return nil, err
	}

	post, err := u.repo.GetPost(ctx, postID)
	if err != nil {
		return nil, err
	}
	if !post.AllowComments {
		return nil, errors.New("comments are not allowed for this post")
	}

	var parentID int64
	if args["parent_id"] != nil {
		tempID, err := strconv.ParseInt(args["parent_id"].(string), 10, 64)
		if err != nil {
			return nil, err
		}
		parentID = tempID
	}

	if parentID != 0 && !u.parentIsValid(ctx, parentID, postID) {
		return nil, errors.New("invalid parentId for this comment")
	}

	authorID, err := strconv.ParseInt(args["author_id"].(string), 10, 64)
	if err != nil {
		return nil, err
	}

	comment := &models.Comment{
		PostID:   postID,
		ParentID: parentID,
		Content:  args["content"].(string),
		AuthorID: authorID,
	}

	err = u.repo.CreateComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// GetComment get comment by id
func (u *commentsUsecase) GetComment(ctx context.Context, args map[string]interface{}) (*models.Comment, error) {
	id, err := strconv.ParseInt(args["id"].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	comments, err := u.repo.GetComment(ctx, id)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// GetComments get comments
func (u *commentsUsecase) GetComments(ctx context.Context, args map[string]interface{}) ([]*models.Comment, error) {
	arg, ok := args["post_id"]
	if !ok {
		arg = args["id"]
	}

	postID, err := strconv.ParseInt(arg.(string), 10, 64)
	if err != nil {
		return nil, err
	}

	limit, ok := args["limit"].(int)
	if !ok {
		limit = defaultLimit
	}

	offset, ok := args["offset"].(int)
	if !ok {
		offset = defaultOffset
	}

	comments, err := u.repo.GetCommentsByPostID(ctx, postID, limit, offset)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (u *commentsUsecase) parentIsValid(ctx context.Context, parentID, postID int64) bool {
	parentComment, err := u.repo.GetComment(ctx, parentID)
	if err != nil {
		log.Printf("failed to get parent comment by id: %v", err)
		return false
	}

	if parentComment.PostID != postID {
		return false
	}

	return true
}
