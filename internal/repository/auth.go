package repository

import (
	"context"

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
