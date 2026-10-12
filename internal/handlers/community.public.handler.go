package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

// GetPublicCommunities godoc
// @Summary Get Public Communities
// @Description List active communities for the Explore page
// @Tags communities
// @Produce json
// @Success 200 {object} dto.CommunityListResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/communities/public [get]
func (h *CommunityHandler) GetPublicCommunities(c *gin.Context) {
	communities, err := h.Service.GetPublicCommunities(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to load communities",
		})
		return
	}

	c.JSON(http.StatusOK, dto.CommunityListResponse{
		Data: communities,
	})
}
