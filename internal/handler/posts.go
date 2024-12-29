package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) newPost(c *gin.Context) {
	// id, err := getUserId(c)
	// if err != nil {
	// 	models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	// 	return
	// }

	// var req models.NewPostRequest
	// if err := c.BindJSON(&req); err != nil {
	// 	models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// post, err := h.service.NewPost(id, req)
	// if err != nil {
	// 	if err == errors.ErrInvalidTag ||
	// 		err == errors.ErrInvalidContent {
	// 		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	// 		return
	// 	}

	// 	models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusCreated, post)
}
