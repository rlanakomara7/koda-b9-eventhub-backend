package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

// SaveCommunity godoc
// @Summary Save Community
// @Description Save a community to the authenticated user's favorites
// @Tags communities
// @Produce json
// @Security BearerToken
// @Param id path int true "Community ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/communities/{id}/save [post]
func (h *CommunityMemberHandler) SaveCommunity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid community id"})
		return
	}

	err = h.Service.SaveCommunity(uint(id), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save community"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "community saved"})
}

// UnsaveCommunity godoc
// @Summary Unsave Community
// @Description Remove a community from the authenticated user's favorites
// @Tags communities
// @Produce json
// @Security BearerToken
// @Param id path int true "Community ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/communities/{id}/save [delete]
func (h *CommunityMemberHandler) UnsaveCommunity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid community id"})
		return
	}

	err = h.Service.UnsaveCommunity(uint(id), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to unsave community"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "community unsaved"})
}

// GetSavedCommunities godoc
// @Summary Get Saved Communities
// @Description Get communities saved by the authenticated user
// @Tags communities
// @Produce json
// @Security BearerToken
// @Success 200 {object} dto.CommunityListResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/communities/saved [get]
func (h *CommunityMemberHandler) GetSavedCommunities(c *gin.Context) {
	communities, err := h.Service.GetSavedCommunities(c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get saved communities"})
		return
	}

	c.JSON(http.StatusOK, dto.CommunityListResponse{
		Data: communities,
	})
}
