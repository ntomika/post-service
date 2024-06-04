package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"post-service/internal/models"
)

func (r *postgresRepo) CreatePost(ctx context.Context, post *models.Post) error {
	query := `INSERT INTO posts (title, content, author_id, allow_comments) 
              VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := r.db.QueryRow(ctx, query, post.Title, post.Content, post.AuthorID, post.AllowComments).Scan(&post.ID, &post.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresRepo) GetPost(ctx context.Context, id int64) (*models.Post, error) {
	query := `SELECT id, title, content, author_id, allow_comments, created_at FROM posts WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	post := &models.Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.AllowComments, &post.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return post, nil
}

func (r *postgresRepo) GetPosts(ctx context.Context) ([]*models.Post, error) {
	query := `SELECT id, title, content, author_id, allow_comments, created_at FROM posts`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.AllowComments, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postgresRepo) CreateComment(ctx context.Context, comment *models.Comment) error {
	query := `INSERT INTO comments (post_id, parent_id, content, author_id) 
              VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := r.db.QueryRow(ctx, query, comment.PostID, comment.ParentID, comment.Content, comment.AuthorID).Scan(&comment.ID, &comment.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresRepo) GetCommentsByPostID(
	ctx context.Context,
	postID int64,
	limit, offset int,
) ([]*models.Comment, error) {
	query := `
        SELECT id, post_id, parent_id, content, author_id, created_at
        FROM comments
        WHERE post_id = $1
        OFFSET $2
        LIMIT $3
    `

	rows, err := r.db.Query(ctx, query, postID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.Content, &comment.AuthorID, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *postgresRepo) GetComment(ctx context.Context, id int64) (*models.Comment, error) {
	query := `SELECT id, post_id, parent_id, content, author_id, created_at FROM comments WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	comment := &models.Comment{}
	err := row.Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.Content, &comment.AuthorID, &comment.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}

	return comment, nil
}
