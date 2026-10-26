package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

// GetEventsPage godoc
// @Summary Get paginated events
// @Description Get events with pagination and filters
// @Tags events
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Param search query string false "Event title"
// @Param format query string false "online or in_person"
// @Param category query string false "Event category"
// @Success 200 {object} dto.EventPaginationResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/events [get]
func (h *EventHandler) GetEventsPage(c *gin.Context) {
	page := 1
	limit := 6

	if raw, exists := c.GetQuery("page"); exists {
		parsed, err := strconv.Atoi(raw)

		if err != nil || parsed < 1 || parsed > 1000000 {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: "page must be between 1 and 1000000",
			})
			return
		}

		page = parsed
	}

	if raw, exists := c.GetQuery("limit"); exists {
		parsed, err := strconv.Atoi(raw)

		if err != nil || parsed < 1 || parsed > 100 {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: "limit must be between 1 and 100",
			})
			return
		}

		limit = parsed
	}

	search := strings.TrimSpace(c.Query("search"))
	format := strings.TrimSpace(c.Query("format"))
	category := strings.TrimSpace(c.Query("category"))

	if format != "" &&
		format != "online" &&
		format != "in_person" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "format must be online or in_person",
		})
		return
	}

	result, err := h.Service.GetEventsPaginated(
		c.Request.Context(),
		search,
		format,
		category,
		page,
		limit,
		c.Request.URL.Path,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "failed to load events",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
