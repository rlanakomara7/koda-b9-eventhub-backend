package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
)

func InitMainRouter(db *pgxpool.Pool) *gin.Engine {

	router := gin.Default()

	userRepository := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepository)

	authHandler := handlers.NewAuthHandler(authService)

	AuthRoutes(router, authHandler)

	return router
}
