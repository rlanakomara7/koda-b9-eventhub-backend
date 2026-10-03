package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
}

func NewUserHandler(userRepository repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: &userRepository,
	}
}

func (h *UserHandler) Profile(c *gin.Context) {

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unautorized",
		})
		return
	}

	profile, err := h.userRepository.GetProfile(userID)

	if err != nil {
		log.Println("GET PROFILE ERROR:", err) // ✅ TAMBAH

		c.JSON(http.StatusInternalServerError, gin.H{ // 🔄 UBAH
			"message": "failed to get profile", // 🔄 UBAH
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "get profile success",
		"data":    profile,
	})
}
