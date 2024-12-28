package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Country struct {
	Name   string `json:"name"`
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
	Region string `json:"region"`
}

type RegisterRequest struct {
	Login       string `json:"login" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CountryCode string `json:"countryCode" binding:"required"`
	IsPublic    bool   `json:"isPublic" binding:"required"`
	Phone       string `json:"phone,omitempty"`
	Image       string `json:"image,omitempty"`
}

type SignInRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type UserProfile struct {
	Login       string `json:"login" binding:"required"`
	Email       string `json:"email" binding:"required"`
	CountryCode string `json:"countryCode" binding:"required"`
	IsPublic    bool   `json:"isPublic" binding:"required"`
	Phone       string `json:"phone,omitempty"`
	Image       string `json:"image,omitempty"`
}

type UpdateProfileRequest struct {
	CountryCode *string `json:"countryCode"`
	IsPublic    *bool   `json:"isPublic"`
	Phone       *string `json:"phone"`
	Image       *string `json:"image"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type FriendInfo struct {
	Login   string    `json:"login"`
	AddedAt time.Time `json:"addedAt"`
}

type LoginRequest struct {
	Login string `json:"login" binding:"required"`
}
