package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/middleware"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

func CommunityRoutes(router *gin.Engine, handler *handlers.CommunityHandler, tokenRepo *repositories.TokenRepository) {

	community := router.Group("/api/communities")

	community.Use(middleware.AuthMiddleware(tokenRepo))

	{
		community.GET("", handler.GetCommunities)
	}
}
