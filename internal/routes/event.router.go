package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/middleware"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

func EventRoutes(
	router *gin.Engine,
	handler *handlers.EventHandler,
	memberHandler *handlers.EventMemberHandler,
	tokenRepo *repositories.TokenRepository,
) {

	event := router.Group("/api/events")

	// public
	event.GET("", handler.GetEvents)
	event.GET("/upcoming", handler.GetUpcomingEvents)
	event.GET("/:id", handler.GetEventDetail)

	// protected
	protected := event.Group("")

	protected.Use(middleware.AuthMiddleware(tokenRepo))

	{
		protected.POST("/:id/join", memberHandler.JoinEvent)

		protected.DELETE("/:id/leave", memberHandler.LeaveEvent)

		protected.GET("/my", handler.GetMyEvents)
	}
}
