package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (h *Handler) addFriend(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req models.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.AddFriend(id, req.Login); err != nil {
		if err == errors.ErrLoginDoesNotExist {
			models.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) removeFriend(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req models.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.RemoveFriend(id, req.Login); err != nil {
		if err == errors.ErrLoginDoesNotExist {
			models.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) listFriends(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	limit_ := c.Query("limit")
	offset_ := c.Query("offset")

	limit, err := strconv.Atoi(limit_)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offset, err := strconv.Atoi(offset_)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	friends, err := h.service.ListFriends(id, limit, offset)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, friends)
}