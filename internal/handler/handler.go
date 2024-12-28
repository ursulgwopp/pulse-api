package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

type Service interface {
	ListCountries(regions []string) ([]models.Country, error)
	GetCountryByAlpha2(alpha2 string) (models.Country, error)

	Register(req models.RegisterRequest) (models.UserProfile, error)
	SignIn(req models.SignInRequest) (string, error)
	ValidateToken(token string) error

	GetProfile(id int) (models.UserProfile, error)
	UpdateProfile(id int, req models.UpdateProfileRequest) (models.UserProfile, error)
	UpdatePassword(id int, req models.UpdatePasswordRequest) error
	GetProfileByLogin(id int, login string) (models.UserProfile, error)

	AddFriend(id int, login string) error
	RemoveFriend(id int, login string) error
	ListFriends(id int, limit int, offset int) ([]models.FriendInfo, error)
}

type Handler struct {
	service Service
}

func NewTransport(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})

		api.GET("/countries", h.listCountries)
		api.GET("/countries/:alpha2", h.getCountryByAlpha2)

		api.POST("/auth/register", h.register)
		api.POST("/auth/sign-in", h.signIn)

		api.GET("/me/profile", h.userIdentity, h.getProfile)
		api.PATCH("/me/profile", h.userIdentity, h.updateProfile)
		api.POST("/me/updatePassword", h.userIdentity, h.updatePassword)

		api.GET("/profiles/:login", h.userIdentity, h.getProfileByLogin)

		api.POST("/friends/add", h.userIdentity, h.addFriend)
		api.POST("/friends/remove", h.userIdentity, h.removeFriend)
		api.GET("/friends", h.userIdentity, h.listFriends)

		// api.POST("")
	}

	return router
}
