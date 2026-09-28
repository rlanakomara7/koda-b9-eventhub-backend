package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

// GET EVENT LIST

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

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": events,
		},
	)

}

// get event detail
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
		c.JSON(404, gin.H{
			"message": "event not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": event,
	})
}

// join event
func (h *EventMemberHandler) JoinEvent(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid event id",
		})
		return
	}

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	uid := uint(userID.(float64))

	err = h.Service.JoinEvent(
		uint(id),
		uid,
	)

	if err != nil {

		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "join event success",
	})
}
