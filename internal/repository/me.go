package repository

import (
	"context"

	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (r *PostgresRepository) GetProfile(id int) (models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var userProfile models.UserProfile
	query := `SELECT login, email, country_code, is_public, phone, image FROM users WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&userProfile.Login, &userProfile.Email, &userProfile.CountryCode, &userProfile.IsPublic, &userProfile.Phone, &userProfile.Image); err != nil {
		return models.UserProfile{}, err
	}

	return userProfile, nil
}

func (r *PostgresRepository) UpdateProfile(id int, req models.UpdateProfileRequest) (models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	if req.CountryCode != nil {
		_, err := r.db.ExecContext(ctx, "UPDATE users SET country_code = $1 WHERE id = $2", *req.CountryCode, id)
		if err != nil {
			return models.UserProfile{}, err
		}
	}

	if req.IsPublic != nil {
		_, err := r.db.ExecContext(ctx, "UPDATE users SET is_public = $1 WHERE id = $2", *req.IsPublic, id)
		if err != nil {
			return models.UserProfile{}, err
		}
	}

	if req.Phone != nil {
		_, err := r.db.ExecContext(ctx, "UPDATE users SET phone = $1 WHERE id = $2", *req.Phone, id)
		if err != nil {
			return models.UserProfile{}, err
		}
	}

	if req.Image != nil {
		_, err := r.db.ExecContext(ctx, "UPDATE users SET image = $1 WHERE id = $2", *req.Image, id)
		if err != nil {
			return models.UserProfile{}, err
		}
	}

	userProfile, err := r.GetProfile(id)
	if err != nil {
		return models.UserProfile{}, err
	}

	return userProfile, nil
}

func (r *PostgresRepository) UpdatePassword(id int, req models.UpdatePasswordRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var oldPassword string
	query := `SELECT hash_password FROM users WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&oldPassword); err != nil {
		return err
	}

	if oldPassword != req.OldPassword {
		return errors.ErrInvalidOldPassword
	}

	query = `UPDATE users SET hash_password = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, req.NewPassword, id)

	return err
}
