package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
)

type CommunityHandler struct {
	Service *services.CommunityService
}

func NewCommunityHandler(services *services.CommunityService) *CommunityHandler {
	return &CommunityHandler{
		Service: services,
	}
}

// Get Communities
//
// @Summary      Get Communities
// @Description  Get community list with search and category filter
// @Tags         communities
// @Produce      json
// @Security	 BearerToken
// @Param        search query string false "Search community by name"
// @Param        category query string false "Filter community by category"
// @Success      200 {object} dto.CommunityListResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/communities [get]
func (h *CommunityHandler) GetCommunities(c *gin.Context) {

	search := c.Query("search")
	category := c.Query("category")

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(
			http.StatusUnauthorized,
			dto.ErrorResponse{
				Message: "unauthorized",
			},
		)
		return

		communities, err := h.Service.GetCommunities(
			search,
			category,
			userID,
		)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dto.ErrorResponse{
					Message: err.Error(),
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			dto.CommunityListResponse{
				Data: communities,
			},
		)
	}
}
