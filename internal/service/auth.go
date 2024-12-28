package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}

func validateLogin(s *Service, login string) error {
	if len(login) > 30 {
		return errors.ErrInvalidLogin
	}

	if matched, _ := regexp.MatchString(`[a-zA-Z0-9-]+`, login); !matched {
		return errors.ErrInvalidLogin
	}

	exists, err := s.repo.CheckLoginExists(login)
	if err != nil {
		return err
	}

	if exists {
		return errors.ErrLoginExists
	}

	return nil
}

func validateEmail(email string) error {
	if len(email) < 1 || len(email) > 50 {
		return errors.ErrInvalidEmail
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email); !matched {
		return errors.ErrInvalidEmail
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 || len(password) > 100 {
		return errors.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
		return errors.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
		return errors.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
		return errors.ErrInvalidPassword
	}

	return nil
}

func validateCountryCode(s *Service, alpha2 string) error {
	exists, err := s.repo.CheckCountryCodeExists(alpha2)
	if err != nil {
		return err
	}

	if !exists {
		return errors.ErrInvalidCountryCode
	}

	return nil
}

func validatePhone(s *Service, phone string) error {
	if phone == "" {
		return nil
	}

	exists, err := s.repo.CheckPhoneExists(phone)
	if err != nil {
		return err
	}

	if exists {
		return errors.ErrPhoneExists
	}

	if len(phone) > 20 {
		return errors.ErrInvalidPhone
	}

	if matched, _ := regexp.MatchString(`^\+?[1-9]\d{0,2}[-.\s]?(\(?\d{1,4}?\)?[-.\s]?)?\d{1,4}[-.\s]?\d{1,4}[-.\s]?\d{1,9}$`, phone); !matched {
		return errors.ErrInvalidPhone
	}

	return nil
}

func validateImage(image string) error {
	if image == "" {
		return nil
	}

	if len(image) > 200 {
		return errors.ErrInvalidImage
	}

	return nil
}

func (s *Service) Register(req models.RegisterRequest) (models.UserProfile, error) {
	if err := validateLogin(s, req.Login); err != nil {
		return models.UserProfile{}, err
	}

	if err := validateEmail(req.Email); err != nil {
		return models.UserProfile{}, err
	}

	if err := validatePassword(req.Password); err != nil {
		return models.UserProfile{}, err
	}

	if err := validateCountryCode(s, req.CountryCode); err != nil {
		return models.UserProfile{}, err
	}

	if err := validatePhone(s, req.Phone); err != nil {
		return models.UserProfile{}, err
	}

	if err := validateImage(req.Image); err != nil {
		return models.UserProfile{}, err
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.Register(req)
}

// probably should completely remove tokenInfo struct, instead return id as int from repository layer
// think on later
func (s *Service) SignIn(req models.SignInRequest) (string, error) {
	req.Password = generatePasswordHash(req.Password)

	tokenClaims, err := s.repo.SignIn(req)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", errors.ErrInvalidUsernameOrPassword
		}

		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: tokenClaims.UserId,
	})

	token, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	if err := s.repo.AddToken(tokenClaims.UserId, token); err != nil {
		return "", err
	}

	return token, nil
}
