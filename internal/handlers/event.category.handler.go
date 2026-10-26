package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

// GetEventCategories godoc
// @Summary Get event categories
// @Description Get all available event categories
// @Tags events
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/events/categories [get]
func (h *EventHandler) GetEventCategories(c *gin.Context) {
	categories, err := h.Service.EventRepo.GetEventCategories(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "failed to load event categories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}
