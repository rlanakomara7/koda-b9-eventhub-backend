package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/middleware"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

func UserRoutes(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
	tokenRepo *repositories.TokenRepository,
) {

	profile := router.Group("/api/profile")

	profile.Use(middleware.AuthMiddleware(tokenRepo))

	{
		profile.GET("", userHandler.Profile)

		profile.PUT("", userHandler.UpdateProfile)
	}
}
