package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/middleware"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

func NotificationRoutes(
	router *gin.Engine,
	handler *handlers.NotificationHandler,
	tokenRepo *repositories.TokenRepository,
) {
	notification := router.Group("/api/notifications")

	notification.Use(middleware.AuthMiddleware(tokenRepo))

	notification.GET("", handler.GetNotifications)
	notification.PATCH("/read-all", handler.MarkAllRead)
	notification.PATCH("/:id/read", handler.MarkRead)
}
