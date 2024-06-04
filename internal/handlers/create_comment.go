package handlers

import (
	"errors"
	"github.com/graphql-go/graphql"
)

// CreateComment creates a new comment for post
func (r *Server) CreateComment(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context

	comment, err := r.commentsUsecase.CreateComment(ctx, p.Args)
	if err != nil {
		return nil, errors.Join(errors.New("failed to create new comment"), err)
	}

	return comment, nil
}
