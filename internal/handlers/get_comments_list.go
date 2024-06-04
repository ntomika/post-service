package handlers

import (
	"errors"
	"github.com/graphql-go/graphql"
)

// GetCommentList get list of comments for particular post
func (r *Server) GetCommentList(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context

	comments, err := r.commentsUsecase.GetComments(ctx, p.Args)
	if err != nil {
		return nil, errors.Join(errors.New("failed to get comments"), err)
	}

	return comments, nil
}
