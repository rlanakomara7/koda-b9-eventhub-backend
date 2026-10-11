package handlers

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/utils"
)

// UpdateEvent godoc
// @Summary Update Event
// @Description Update an event owned by the authenticated organizer. Send all form fields; banner is optional.
// @Tags events
// @Accept mpfd
// @Produce json
// @Security BearerToken
// @Param id path int true "Event ID"
// @Param title formData string true "Event title"
// @Param description formData string false "Description"
// @Param category formData string true "Existing category name"
// @Param community_id formData int false "Community ID"
// @Param event_date formData string true "YYYY-MM-DD"
// @Param start_time formData string true "HH:MM"
// @Param end_time formData string true "HH:MM"
// @Param format formData string true "online or in_person"
// @Param location formData string false "Location"
// @Param capacity formData int true "Capacity"
// @Param speakers formData []string false "Speaker names" collectionFormat(multi)
// @Param banner formData file false "New banner (optional)"
// @Success 200 {object} dto.UpdateEventResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/events/{id} [patch]
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event ID",
		})
		return
	}

	var req dto.UpdateEventRequest

	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event form data",
		})
		return
	}

	req.BannerURL, err = utils.SaveUploadedFile(c, "banner")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.Service.UpdateEvent(userID, uint(id), req)

	if err != nil {
		if req.BannerURL != "" {
			_ = os.Remove("." + req.BannerURL)
		}

		switch {
		case errors.Is(err, repositories.ErrUpdateEventNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": "event not found",
			})

		case errors.Is(err, repositories.ErrUpdateEventForbidden),
			errors.Is(err, repositories.ErrNotOrganizer):
			c.JSON(http.StatusForbidden, gin.H{
				"message": "you do not have permission to edit this event",
			})

		case errors.Is(err, repositories.ErrCategoryNotFound),
			errors.Is(err, repositories.ErrCommunityNotFound),
			errors.Is(err, services.ErrInvalidEventData):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to update event",
			})
		}
		return
	}

	c.JSON(http.StatusOK, dto.UpdateEventResponse{
		Message: "event updated successfully",
		EventID: uint(id),
	})
}
