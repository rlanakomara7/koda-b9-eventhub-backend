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

	community.POST("/:id/join", middleware.AuthMiddleware(tokenRepo), memberHandler.JoinCommunity)

	community.DELETE("/:id/leave", middleware.AuthMiddleware(tokenRepo), memberHandler.LeaveCommunity)
}
