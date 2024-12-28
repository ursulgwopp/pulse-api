package repository

import (
	"context"

	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (r *PostgresRepository) AddToken(id int, token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `INSERT INTO tokens (user_id, token, is_valid) VALUES ($1, $2, TRUE)`
	_, err := r.db.ExecContext(ctx, query, id, token)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) ValidateToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var is_valid bool
	query := `SELECT is_valid FROM tokens WHERE token = $1`
	if err := r.db.QueryRowContext(ctx, query, token).Scan(&is_valid); err != nil {
		return err
	}

	if !is_valid {
		return errors.ErrInvalidToken
	}

	return nil
}

func (r *PostgresRepository) KillTokens(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE tokens SET is_valid = FALSE WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
