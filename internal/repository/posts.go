package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (r *PostgresRepository) NewPost(login string, req models.NewPostRequest) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var post models.Post
	var likesCount sql.NullInt32
	var dislikesCount sql.NullInt32

	query := `INSERT INTO posts (id, content, author, tags) VALUES ($1, $2, $3, $4) RETURNING id, content, author, tags, created_at, array_length(likes_count, 1), array_length(dislikes_count, 1)`
	if err := r.db.QueryRowContext(ctx, query, uuid.New(), req.Content, login, pq.Array(req.Tags)).Scan(&post.Id, &post.Content, &post.Author, pq.Array(&post.Tags), &post.CreatedAt, &likesCount, &dislikesCount); err != nil {
		return models.Post{}, err
	}

	post.LikesCount = likesCount.Int32
	post.DislikesCount = dislikesCount.Int32

	return post, nil
}
