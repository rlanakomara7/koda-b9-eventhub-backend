package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
)

type NotificationHandler struct {
	Service *services.NotificationService
}

func NewNotificationHandler(
	service *services.NotificationService,
) *NotificationHandler {
	return &NotificationHandler{Service: service}
}

// GetNotifications godoc
// @Summary Get Notifications
// @Description Get notifications for the authenticated user
// @Tags notifications
// @Produce json
// @Security BearerToken
// @Success 200 {object} dto.NotificationListResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/notifications [get]
func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	notifications, err := h.Service.GetNotifications(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to load notifications",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NotificationListResponse{
		Data: notifications,
	})
}

// MarkRead godoc
// @Summary Mark Notification as Read
// @Description Mark one notification as read
// @Tags notifications
// @Produce json
// @Security BearerToken
// @Param id path int true "Notification ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/notifications/{id}/read [patch]
func (h *NotificationHandler) MarkRead(c *gin.Context) {
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
			"message": "invalid notification id",
		})
		return
	}

	err = h.Service.MarkRead(
		c.Request.Context(),
		userID,
		uint(id),
	)

	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "notification not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to mark notification as read",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "notification marked as read",
	})
}

// MarkAllRead godoc
// @Summary Mark All Notifications as Read
// @Description Mark all notifications of the authenticated user as read
// @Tags notifications
// @Produce json
// @Security BearerToken
// @Success 200 {object} dto.Response
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/notifications/read-all [patch]
func (h *NotificationHandler) MarkAllRead(c *gin.Context) {
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	err := h.Service.MarkAllRead(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to mark all notifications as read",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all notifications marked as read",
	})
}
