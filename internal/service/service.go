package service

import "github.com/ursulgwopp/pulse-api/internal/models"

type Repository interface {
	ListCountries(regions []string) ([]models.Country, error)
	GetCountryByAlpha2(alpha2 string) (models.Country, error)

	Register(req models.RegisterRequest) (models.UserProfile, error)
	SignIn(req models.SignInRequest) (models.TokenClaims, error)
	AddToken(id int, token string) error
	ValidateToken(token string) error
	KillTokens(id int) error

	GetProfile(id int) (models.UserProfile, error)
	UpdateProfile(id int, req models.UpdateProfileRequest) (models.UserProfile, error)
	UpdatePassword(id int, req models.UpdatePasswordRequest) error

	AddFriend(id int, login string) error
	RemoveFriend(id int, login string) error
	ListFriends(id int, limit int, offset int) ([]models.FriendInfo, error)

	CheckLoginExists(login string) (bool, error)
	CheckCountryCodeExists(alpha2 string) (bool, error)
	CheckPhoneExists(phone string) (bool, error)
	// CheckUserIdExists(id int) (bool, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
