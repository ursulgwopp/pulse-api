package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		models.NewErrorResponse(c, http.StatusUnauthorized, errors.ErrEmptyAuthHeader.Error())
		return
	}

	token, _ := strings.CutPrefix(header, "Bearer ")

	if err := h.service.ValidateToken(token); err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	userId, err := parseToken(token)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("token", token)
	c.Set("user_id", userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("user_id")
	if !ok {
		return 0, errors.ErrUserIdNotFound
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.ErrInvalidIdType
	}

	return idInt, nil
}

func parseToken(token string) (int, error) {
	token_, err := jwt.ParseWithClaims(token, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidSigningMethod
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return -1, err
	}

	claims, ok := token_.Claims.(*models.TokenClaims)
	if !ok {
		return -1, errors.ErrInvalidTokenClaims
	}

	return claims.UserId, nil
}
