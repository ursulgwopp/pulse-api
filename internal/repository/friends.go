package repository

import (
	"context"
	"time"

	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (r *PostgresRepository) AddFriend(id int, login string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var userLogin string
	query := `SELECT login FROM users WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&userLogin); err != nil {
		return err
	}

	if userLogin == login {
		return nil
	}

	var exists bool
	query = `SELECT EXISTS (SELECT 1 FROM friends WHERE user_id = $1 AND EXISTS (SELECT 1 FROM UNNEST(friends_info) AS fi WHERE fi.login = $2))`
	if err := r.db.QueryRowContext(ctx, query, id, login).Scan(&exists); err != nil {
		return err
	}

	if exists {
		return nil
	}

	query = `UPDATE friends SET friends_info = ARRAY_APPEND(friends_info, ROW($1, $2::TIMESTAMP)::friend_info) WHERE user_id = $3`
	_, err := r.db.ExecContext(ctx, query, login, time.Now(), id)

	return err
}

func (r *PostgresRepository) RemoveFriend(id int, login string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE friends SET friends_info = (SELECT ARRAY_AGG(fi) FROM UNNEST(friends_info) AS fi WHERE fi.login <> $1) WHERE user_id = $2`
	_, err := r.db.ExecContext(ctx, query, login, id)

	return err
}

func (r *PostgresRepository) ListFriends(id int, limit int, offset int) ([]models.FriendInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var friends []models.FriendInfo
	// query := `SELECT (UNNEST(friends_info)).* FROM friends WHERE user_id = $1`
	query := `WITH unnested_friends AS (SELECT (UNNEST(friends_info)).* AS friend_info FROM friends WHERE user_id = $1) SELECT * FROM unnested_friends LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, id, limit, offset)
	if err != nil {
		return []models.FriendInfo{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var friend models.FriendInfo
		if err := rows.Scan(&friend.Login, &friend.AddedAt); err != nil {
			return []models.FriendInfo{}, err
		}

		friends = append(friends, friend)
	}

	if rows.Err() != nil {
		return []models.FriendInfo{}, err
	}

	return friends, nil
}
