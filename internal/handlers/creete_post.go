package handlers

import (
	"errors"
	"github.com/graphql-go/graphql"
)

// CreatePost creates a new post
func (r *Server) CreatePost(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context

	post, err := r.postsUsecase.CreatePost(ctx, p.Args)
	if err != nil {
		return nil, errors.Join(errors.New("failed to create new post"), err)
	}

	return post, nil
}
