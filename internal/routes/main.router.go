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

	//repostiory
	userRepository := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)
	passwordRepo := repositories.NewPasswordRepository(db)

	//service
	authService := services.NewAuthService(userRepository)
	passwordService := services.NewPasswordService(userRepository, passwordRepo)

	//handler
	authHandler := handlers.NewAuthHandler(authService, tokenRepo, passwordService)

	//event
	eventRepo := repositories.NewEventRepository(db)
	eventService := services.NewEventService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)

	//event member
	eventMemberRepo := repositories.NewEventMemberRepository(db)

	eventMemberService := services.NewEventMemberService(eventMemberRepo)

	eventMemberHandler := handlers.NewEventMemberHandler(eventMemberService)

	EventRoutes(router, eventHandler, eventMemberHandler, tokenRepo)
	//route
	AuthRoutes(router, authHandler)

	return router
}
