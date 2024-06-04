package handlers

import (
	"errors"
	"github.com/graphql-go/graphql"
)

// GetPostList gets list of posts
func (r *Server) GetPostList(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context

	posts, err := r.postsUsecase.GetPosts(ctx)
	if err != nil {
		return nil, errors.Join(errors.New("failed to get posts"), err)
	}

	return posts, nil
}
