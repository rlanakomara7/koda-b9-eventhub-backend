package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetJoinedEvents godoc
// @Summary Get events joined by current user
// @Tags events
// @Produce json
// @Security BearerToken
// @Success 200 {object} dto.EventCardListResponse
// @Router /api/events/joined [get]
func (h *EventMemberHandler) GetJoinedEvents(c *gin.Context) {
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	events, err := h.Service.GetJoinedEvents(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to load joined events",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}
