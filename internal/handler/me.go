package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (h *Handler) getProfile(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, errors.ErrEmptyAlpha2.Error())
		return
	}

	userProfile, err := h.service.GetProfile(id)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, errors.ErrEmptyAlpha2.Error())
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func (h *Handler) updateProfile(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req models.UpdateProfileRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userProfile, err := h.service.UpdateProfile(id, req)
	if err != nil {
		if err == errors.ErrInvalidCountryCode ||
			err == errors.ErrInvalidPhone ||
			err == errors.ErrInvalidImage {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err == errors.ErrPhoneExists {
			models.NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func (h *Handler) updatePassword(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req models.UpdatePasswordRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdatePassword(id, req); err != nil {
		if err == errors.ErrInvalidPassword {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err == errors.ErrInvalidOldPassword {
			models.NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}
