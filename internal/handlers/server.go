package handlers

import (
	"post-service/internal/repository"
	"post-service/internal/usecases"
)

type Server struct {
	postsUsecase    usecases.Posts
	commentsUsecase usecases.Comments
}

func NewServer(repo repository.Repository) *Server {
	return &Server{
		postsUsecase:    usecases.NewPostsUsecase(repo),
		commentsUsecase: usecases.NewComentsUsecase(repo),
	}
}
