package handlers

import (
	"errors"
	"github.com/graphql-go/graphql"
)

// GetPost get post by ID
func (r *Server) GetPost(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context

	post, err := r.postsUsecase.GetPost(ctx, p.Args)
	if err != nil {
		return nil, errors.Join(errors.New("failed to get post"), err)
	}

	comments, err := r.commentsUsecase.GetComments(ctx, p.Args)
	if err != nil {
		return nil, errors.Join(errors.New("failed to get comments for post"), err)
	}

	post.Comments = comments

	return post, nil
}
