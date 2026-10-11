package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/utils"
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

// Update Profile
//
// @Summary      Update Profile
// @Description  Update profile of currently logged in user
// @Tags         user
// @Accept       mpfd
// @Produce      json
// @Security     BearerToken
// @Param        name formData string true "Name"
// @Param        location formData string false "Location"
// @Param        bio formData string false "Bio"
// @Param        avatar formData file false "Avatar"
// @Success      200 {object} dto.ProfileSuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      422 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/profile [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(
			http.StatusUnauthorized,
			dto.ErrorResponse{
				Message: "unauthorized",
			},
		)
		return
	}

	var req dto.UpdateProfileRequest

	// multipart/form-data + Gin Validator
	err := c.ShouldBindWith(
		&req,
		binding.FormMultipart,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
			},
		)
		return
	}

	//cekstring
	req.Name = strings.TrimSpace(req.Name)
	req.Location = strings.TrimSpace(req.Location)
	req.Bio = strings.TrimSpace(req.Bio)

	// upload avatar
	avatarURL, err := utils.SaveUploadedFile(
		c,
		"avatar",
	)

	if err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			dto.ErrorResponse{
				Message: err.Error(),
			},
		)
		return
	}

	// kalau user tidak upload avatar baru
	if avatarURL == "" {

		currentProfile, err := h.userRepository.GetProfile(
			userID,
		)

		if err != nil {

			fmt.Println(
				"GET PROFILE ERROR:",
				err,
			)

			c.JSON(
				http.StatusInternalServerError,
				dto.ErrorResponse{
					Message: "failed to get profile",
				},
			)
			return
		}

		// pakai avatar lama
		req.AvatarURL = currentProfile.AvatarURL

	} else {

		// pakai avatar baru
		req.AvatarURL = avatarURL
	}

	//  update database
	err = h.userRepository.UpdateProfile(
		userID,
		req,
	)

	if err != nil {

		fmt.Println(
			"UPDATE PROFILE ERROR:",
			err,
		)

		// supaya kita tahu query mana yang salah
		c.JSON(
			http.StatusInternalServerError,
			dto.ErrorResponse{
				Message: err.Error(),
			},
		)
		return
	}

	// ambil profile terbaru
	profile, err := h.userRepository.GetProfile(
		userID,
	)

	if err != nil {

		fmt.Println(
			"GET UPDATED PROFILE ERROR:",
			err,
		)

		c.JSON(
			http.StatusInternalServerError,
			dto.ErrorResponse{
				Message: "failed to get profile",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "update profile success",
			"data":    profile,
		},
	)
}
