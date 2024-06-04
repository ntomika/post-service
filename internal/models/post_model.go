package models

import "time"

type Post struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	AuthorID      int64     `json:"author_id"`
	AllowComments bool      `json:"allow_comments"`
	CreatedAt     time.Time `json:"created_at"`
	Comments      []*Comment
}
