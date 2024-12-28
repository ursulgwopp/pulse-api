package repository

import (
	"context"

	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (r *PostgresRepository) Register(req models.RegisterRequest) (models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var userId int
	var userProfile models.UserProfile
	query := `INSERT INTO users (login, email, hash_password, country_code, is_public, phone, image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, login, email, country_code, is_public, phone, image`
	if err := r.db.QueryRowContext(ctx, query, req.Login, req.Email, req.Password, req.CountryCode, req.IsPublic, req.Phone, req.Image).Scan(&userId, &userProfile.Login, &userProfile.Email, &userProfile.CountryCode, &userProfile.IsPublic, &userProfile.Phone, &userProfile.Image); err != nil {
		return models.UserProfile{}, err
	}

	query = `INSERT INTO friends (user_id) VALUES ($1)`
	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return models.UserProfile{}, err
	}

	return userProfile, nil
}

func (r *PostgresRepository) SignIn(req models.SignInRequest) (models.TokenClaims, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var id int
	query := `SELECT id FROM users WHERE login = $1 AND hash_password = $2`
	if err := r.db.QueryRowContext(ctx, query, req.Login, req.Password).Scan(&id); err != nil {
		return models.TokenClaims{}, err
	}

	return models.TokenClaims{UserId: id}, nil
}

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
