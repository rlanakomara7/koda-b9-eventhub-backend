package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
)

type EventHandler struct {
	Service *services.EventService
}

type EventMemberHandler struct {
	Service *services.EventMemberService
}

func NewEventHandler(
	service *services.EventService,
) *EventHandler {

	return &EventHandler{
		Service: service,
	}
}

func NewEventMemberHandler(
	service *services.EventMemberService,
) *EventMemberHandler {

	return &EventMemberHandler{
		Service: service,
	}
}

// Get Events
//
// @Summary      Get Events
// @Description  Get event list with optional search and format filter
// @Tags         events
// @Produce      json
// @Param        search query string false "Search event by title"
// @Param        format query string false "Filter event format: in_person or online"
// @Success      200 {object} dto.EventListSuccessResponse
// @Failure      404 {object} dto.EventListNotFoundResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/events [get]
func (h *EventHandler) GetEvents(
	c *gin.Context,
) {

	search := c.Query("search")
	format := c.Query("format")

	events, err := h.Service.GetEvents(search, format)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	if search != "" && len(events) == 0 {
		c.JSON(
			http.StatusNotFound,
			dto.EventListNotFoundResponse{
				Message: "event not found",
				Data:    events,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		dto.EventListSuccessResponse{
			Data: events,
		},
	)

}

// Get Event Detail
//
// @Summary      Get Event Detail
// @Description  Get event detail by event ID
// @Tags         events
// @Produce      json
// @Param        id path int true "Event ID"
// @Success      200 {object} dto.EventDetailSuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /api/events/{id} [get]
func (h *EventHandler) GetEventDetail(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}

	event, err := h.Service.GetEventDetail(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
		return
	}

	c.JSON(http.StatusOK,
		dto.EventDetailSuccessResponse{
			Data: event,
		})
}

// Join Event
//
// @Summary      Join Event
// @Description  Join an event for currently logged in user
// @Tags         events
// @Produce      json
// @Security     BearerToken
// @Param        id path int true "Event ID"
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Router       /api/events/{id}/join [post]
func (h *EventMemberHandler) JoinEvent(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid event id",
		})
		return
	}

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized,
			dto.ErrorResponse{
				Message: "unauthorized",
			})
		return
	}

	err = h.Service.JoinEvent(
		uint(id),
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK,
		dto.Response{
			Message: "join event success",
		})
}

// Leave Event
//
// @Summary      Leave Event
// @Description  Leave an event for currently logged in user
// @Tags         events
// @Produce      json
// @Security     BearerToken
// @Param        id path int true "Event ID"
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Router       /api/events/{id}/leave [delete]
func (h *EventMemberHandler) LeaveEvent(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "invalid event id",
			})
		return
	}

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized,
			dto.ErrorResponse{
				Message: "unauthorized",
			})
		return
	}

	err = h.Service.LeaveEvent(
		uint(id),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK,
		dto.Response{
			Message: "leave event success",
		})
}

// Get Upcoming Events
//
// @Summary      Get Upcoming Events
// @Description  Get list of upcoming events
// @Tags         events
// @Produce      json
// @Success      200 {object} dto.EventListSuccessResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/events/upcoming [get]
func (h *EventHandler) GetUpcomingEvents(c *gin.Context) {

	events, err := h.Service.GetUpcomingEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

// Get My Events
//
// @Summary      Get My Events
// @Description  Get events joined by currently logged in user
// @Tags         events
// @Produce      json
// @Security     BearerToken
// @Success      200 {object} dto.EventListSuccessResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/events/my [get]
func (h *EventHandler) GetMyEvents(c *gin.Context) {
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	events, err := h.Service.GetMyEvents(userID)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK,
		dto.EventListSuccessResponse{
			Data: events,
		})
}
