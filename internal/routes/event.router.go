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

	event.GET("", func(c *gin.Context) {
		_, hasPage := c.GetQuery("page")
		_, hasLimit := c.GetQuery("limit")
		_, hasCategory := c.GetQuery("category")

		if hasPage || hasLimit || hasCategory {
			handler.GetEventsPage(c)
			return
		}

		handler.GetEvents(c)
	})

	event.GET("/categories", handler.GetEventCategories)
	event.GET("/upcoming", handler.GetUpcomingEvents)
	event.GET("/:id", handler.GetEventDetail)
	event.GET("/:id/discussions", memberHandler.GetEventDiscussions)

	protected := event.Group("")
	protected.Use(middleware.AuthMiddleware(tokenRepo))

	protected.GET("/my", handler.GetMyEvents)
	protected.GET("/joined", memberHandler.GetJoinedEvents)

	protected.POST("", handler.CreateEvent)
	protected.PATCH("/:id", handler.UpdateEvent)
	protected.DELETE("/:id", handler.DeleteEvent)

	protected.POST("/:id/join", memberHandler.JoinEvent)
	protected.DELETE("/:id/leave", memberHandler.LeaveEvent)
	protected.GET("/saved", memberHandler.GetSavedEvents)
	protected.POST("/:id/save", memberHandler.SaveEvent)
	protected.DELETE("/:id/save", memberHandler.UnsaveEvent)
	protected.POST("/:id/discussions", memberHandler.CreateEventDiscussion)
}
