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
// @Param        search query string false "Search community"
// @Param        category query string false "Filter community by category"
// @Success      200 {object} dto.CommunityListResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/communities [get]
func (h *CommunityHandler) GetCommunities(c *gin.Context) {

	search := c.Query("search")
	category := c.Query("category")
	userID := c.GetUint("user_id")

	communities, err := h.Service.GetCommunities(
		search,
		category,
		userID,
	)

	if err != nil {
		_ = c.Error(err)

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "failed to get communities",
		})
		return
	}

	c.JSON(http.StatusOK, dto.CommunityListResponse{
		Data: communities,
	},
	)
}
