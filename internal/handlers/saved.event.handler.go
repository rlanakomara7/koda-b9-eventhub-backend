package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

// SaveEvent godoc
// @Summary Save Event
// @Description Save an event to the authenticated user's favorites
// @Tags events
// @Produce json
// @Security BearerToken
// @Param id path int true "Event ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/events/{id}/save [post]
func (h *EventMemberHandler) SaveEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	err = h.Service.SaveEvent(uint(id), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event saved"})
}

// UnsaveEvent godoc
// @Summary Unsave Event
// @Description Remove an event from the authenticated user's favorites
// @Tags events
// @Produce json
// @Security BearerToken
// @Param id path int true "Event ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/events/{id}/save [delete]
func (h *EventMemberHandler) UnsaveEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	err = h.Service.UnsaveEvent(uint(id), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to unsave event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event unsaved"})
}

// GetSavedEvents godoc
// @Summary Get Saved Events
// @Description Get events saved by the authenticated user
// @Tags events
// @Produce json
// @Security BearerToken
// @Success 200 {object} dto.EventCardListResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/events/saved [get]
func (h *EventMemberHandler) GetSavedEvents(c *gin.Context) {
	events, err := h.Service.GetSavedEvents(c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get saved events"})
		return
	}

	c.JSON(http.StatusOK, dto.EventCardListResponse{
		Data: events,
	})
}
