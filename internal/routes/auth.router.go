package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/middleware"
)

func AuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {

	auth := router.Group("/api/auth")

	{
		auth.POST("/register", handler.Register)

		auth.POST("/login", handler.Login)
	}

	protected := router.Group("/api")

	protected.Use(middleware.AuthMiddleware())

	{
		protected.GET(
			"/profile",
			handlers.Profile,
		)
	}
}
