package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
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

// Profile
//
// @Summary      Get Profile
// @Description  Get profile of currently logged in user
// @Tags         user
// @Produce      json
// @Security     BearerToken
// @Success      200 {object} dto.ProfileSuccessResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/profile [get]
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

// ✅ TAMBAH
// Update Profile
//
// @Summary      Update Profile
// @Description  Update profile of currently logged in user
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerToken
// @Param        request body dto.UpdateProfileRequest true "Profile data"
// @Success      200 {object} dto.ProfileSuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {

	var request dto.UpdateProfileRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	err := h.userRepository.UpdateProfile(
		userID,
		request,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update profile",
		})
		return
	}

	profile, err := h.userRepository.GetProfile(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update profile success",
		"data":    profile,
	})
}
