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
