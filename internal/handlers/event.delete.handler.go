package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
)

// DeleteEvent godoc
// @Summary Delete Event
// @Description Permanently delete an event owned by the authenticated organizer
// @Tags events
// @Produce json
// @Security BearerToken
// @Param id path int true "Event ID"
// @Success 200 {object} dto.DeleteEventResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/events/{id} [delete]
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	eventID, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || eventID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event ID",
		})
		return
	}

	err = h.Service.DeleteEvent(
		c.Request.Context(),
		userID,
		uint(eventID),
	)

	if err != nil {
		switch {
		case errors.Is(
			err,
			repositories.ErrDeleteEventNotFound,
		):
			c.JSON(http.StatusNotFound, gin.H{
				"message": "event not found",
			})

		case errors.Is(
			err,
			repositories.ErrDeleteEventForbidden,
		):
			c.JSON(http.StatusForbidden, gin.H{
				"message": "you do not have permission to delete this event",
			})

		case errors.Is(
			err,
			services.ErrInvalidEventData,
		):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid event data",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to delete event",
			})
		}
		return
	}

	c.JSON(http.StatusOK, dto.DeleteEventResponse{
		Message: "event deleted successfully",
		EventID: uint(eventID),
	})
}
