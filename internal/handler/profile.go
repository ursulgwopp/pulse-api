package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (h *Handler) getProfileByLogin(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	login := c.Param("login")
	if login == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidLogin.Error())
		return
	}

	userProfile, err := h.service.GetProfileByLogin(id, login)
	if err != nil {
		if err == errors.ErrLoginDoesNotExist ||
			err == errors.ErrAccessDenied {
			models.NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userProfile)
}
