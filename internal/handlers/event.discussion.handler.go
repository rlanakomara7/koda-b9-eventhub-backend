package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (h *EventMemberHandler) GetEventDiscussions(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event id",
		})
		return
	}

	discussions, err := h.Service.GetEventDiscussions(
		c.Request.Context(),
		uint(id),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to load discussions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": discussions,
	})
}

func (h *EventMemberHandler) CreateEventDiscussion(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event id",
		})
		return
	}

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	var body struct {
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	message := strings.TrimSpace(body.Message)

	if message == "" || utf8.RuneCountInString(message) > 1000 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "message must contain 1-1000 characters",
		})
		return
	}

	discussion, err := h.Service.CreateEventDiscussion(
		c.Request.Context(),
		uint(id),
		userID,
		message,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "join this event first or event is unavailable",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create discussion",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "discussion created",
		"data":    discussion,
	})
}
