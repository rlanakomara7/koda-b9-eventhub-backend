package handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/utils"
)

// Create Event
//
// @Summary Create Event
// @Description Create an event as an organizer with optional banner upload
// @Tags events
// @Accept mpfd
// @Produce json
// @Security BearerToken
// @Param title formData string true "Event title"
// @Param description formData string false "Description"
// @Param category formData string true "Existing event category name"
// @Param community_id formData int false "Community ID"
// @Param event_date formData string true "YYYY-MM-DD"
// @Param start_time formData string true "HH:MM"
// @Param end_time formData string true "HH:MM"
// @Param format formData string true "online or in_person"
// @Param location formData string false "Location"
// @Param capacity formData int true "Capacity"
// @Param speakers formData []string false "Speaker names" collectionFormat(multi)
// @Param banner formData file false "Event banner (jpg, jpeg, png, webp, maximum 2MB)"
// @Success 201 {object} dto.CreateEventResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/events [post]
func (h *EventHandler) CreateEvent(c *gin.Context) {
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	var req dto.CreateEventRequest

	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid event form data",
		})
		return
	}

	req.Speakers = c.PostFormArray("speakers")

	var err error

	req.BannerURL, err = utils.SaveUploadedFile(c, "banner")

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	eventID, err := h.Service.CreateEvent(userID, req)

	if err != nil {

		if req.BannerURL != "" {
			_ = os.Remove("." + req.BannerURL)
		}

		switch {
		case errors.Is(err, repositories.ErrNotOrganizer),
			errors.Is(err, repositories.ErrCommunityNotFound),
			errors.Is(err, services.ErrInvalidEventData):
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create event",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, dto.CreateEventResponse{
		Message: "event created successfully",
		EventID: eventID,
	})
}
