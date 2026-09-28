package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
)

type AuthHandler struct {
	Service         *services.AuthService
	TokenRepo       *repositories.TokenRepository
	PasswordService *services.PasswordService
}

func NewAuthHandler(service *services.AuthService, tokenRepo *repositories.TokenRepository, passwordService *services.PasswordService) *AuthHandler {
	return &AuthHandler{
		Service:         service,
		TokenRepo:       tokenRepo,
		PasswordService: passwordService,
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
				"message": err.Error(),
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
			400,
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
				"name":    user.Email,
				"email":   user.Email,
			},
		},
	)
}

// login
func (h *AuthHandler) Login(c *gin.Context) {

	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, token, err := h.Service.Login(
		request.Email,
		request.Password,
	)

	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "login success",

		"token": token,

		"user": gin.H{
			"user_id": user.UserID,
			"name":    user.Name,
			"email":   user.Email,
		},
	})
}

//logout

func (h *AuthHandler) Logout(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	token := strings.TrimSpace(
		strings.TrimPrefix(
			authHeader,
			"Bearer ",
		),
	)

	err := h.TokenRepo.TokenBlackList(token)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "logout failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "logout success",
	})
}

// forgotpassword
func (h *AuthHandler) ForgotPassword(c *gin.Context) {

	var request dto.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	token, err := h.PasswordService.ForgotPassword(
		request.Email,
	)

	if err != nil {

		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{

		"message": "reset token created",

		"token": token,
	})

}

// handler password
func (h *AuthHandler) ResetPassword(c *gin.Context) {

	var request dto.ResetPasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := h.PasswordService.ResetPassword(
		request.Email,
		request.Token,
		request.NewPassword,
	)

	if err != nil {

		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{

		"message": "password updated",
	})

}
