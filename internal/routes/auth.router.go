package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
)

func AuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {

	auth := router.Group("/api/auth")

	{
		auth.POST("/register", handler.Register)
	}

}
