package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/middleware"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

func CommunityRoutes(
	router *gin.Engine,
	handler *handlers.CommunityHandler,
	memberHandler *handlers.CommunityMemberHandler,
	tokenRepo *repositories.TokenRepository,
) {
	community := router.Group("/api/communities")

	community.GET("/public", handler.GetPublicCommunities)

	community.GET("", middleware.OptionalAuthMiddleware(tokenRepo), handler.GetCommunities)

	community.GET("/:id/members", memberHandler.GetMembers)

	protected := community.Group("")
	protected.Use(middleware.AuthMiddleware(tokenRepo))

	protected.POST("/:id/join", memberHandler.JoinCommunity)
	protected.DELETE("/:id/leave", memberHandler.LeaveCommunity)

	protected.GET("/saved", memberHandler.GetSavedCommunities)
	protected.POST("/:id/save", memberHandler.SaveCommunity)
	protected.DELETE("/:id/save", memberHandler.UnsaveCommunity)
}
