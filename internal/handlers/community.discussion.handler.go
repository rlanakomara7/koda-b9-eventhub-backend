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

func (h *CommunityMemberHandler) GetCommunityDiscussions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid community id",
		})
		return
	}

	discussions, err := h.Service.GetCommunityDiscussions(
		c.Request.Context(),
		uint(id),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to load community discussions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": discussions,
	})
}

func (h *CommunityMemberHandler) CreateCommunityDiscussion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid community id",
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

	discussion, err := h.Service.CreateCommunityDiscussion(
		c.Request.Context(),
		uint(id),
		userID,
		message,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "join this community first or community is unavailable",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create community discussion",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "discussion created",
		"data":    discussion,
	})
}
