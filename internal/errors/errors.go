package errors

import "errors"

var (
	ErrInvalidRegion   = errors.New("invalid region")
	ErrEmptyAlpha2     = errors.New("empty alpha2 code")
	ErrCountryNotFound = errors.New("country not found")

	// ErrInvalidRegisterRequest    = errors.New("invalid register request")
	ErrInvalidLogin              = errors.New("invalid login")
	ErrInvalidEmail              = errors.New("invalid email")
	ErrInvalidPassword           = errors.New("invalid password")
	ErrInvalidCountryCode        = errors.New("invalid country code")
	ErrInvalidPhone              = errors.New("invalid phone")
	ErrInvalidImage              = errors.New("invalid image")
	ErrLoginExists               = errors.New("login exists")
	ErrPhoneExists               = errors.New("phone exists")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrInvalidOldPassword        = errors.New("invalid old password")
	ErrLoginDoesNotExist         = errors.New("login does not exist")

	ErrEmptyAuthHeader      = errors.New("empty auth header")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidTokenClaims   = errors.New("token claims are not of type tokenClaims")
	ErrUserIdNotFound       = errors.New("user id not found")
	ErrInvalidIdType        = errors.New("invalid id type")
	ErrInvalidToken         = errors.New("invalid token")
)
