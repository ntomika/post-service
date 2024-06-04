package handlers

import (
	"errors"
	"github.com/graphql-go/graphql"
)

// GetComment get comment by id
func (r *Server) GetComment(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context

	comments, err := r.commentsUsecase.GetComment(ctx, p.Args)
	if err != nil {
		return nil, errors.Join(errors.New("failed to get comment"), err)
	}

	return comments, nil
}
