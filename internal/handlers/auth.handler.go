package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

// Register
func (h *AuthHandler) Register(c *gin.Context) {

	var request dto.RegisterRequest

	//parsing JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid request",
			},
		)
		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err := h.Service.Register(&user)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "register success",
			"user": gin.H{
				"user_id": user.UserID,
				"name":    user.Name,
				"email":   user.Email,
			},
		},
	)
}

// Login
func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid request",
			},
		)
		return
	}

	user, token, err := h.Service.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "login success",
			"token":   token,
			"user": gin.H{
				"user_id": user.UserID,
				"name":    user.Name,
				"email":   user.Email,
			},
		},
	)
}

