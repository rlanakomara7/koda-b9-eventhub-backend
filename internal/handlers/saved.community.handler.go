package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func (h *CommunityMemberHandler) GetSavedCommunities(c *gin.Context) {
	communities, err := h.Service.GetSavedCommunities(c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get saved communities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": communities})
}
