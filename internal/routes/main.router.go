package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/handlers"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/middleware"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitMainRouter(db *pgxpool.Pool, redisClient *redis.Client) *gin.Engine {

	router := gin.Default()

	//cors
	router.Use(middleware.Cors)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//repostiory
	userRepository := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)
	passwordRepo := repositories.NewPasswordRepository(db)

	//service
	authService := services.NewAuthService(userRepository)
	passwordService := services.NewPasswordService(userRepository, passwordRepo)

	//handler
	authHandler := handlers.NewAuthHandler(authService, tokenRepo, passwordService)
	userHandler := handlers.NewUserHandler(*userRepository)

	//event
	eventRepo := repositories.NewEventRepository(db, redisClient)
	eventService := services.NewEventService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)

	//community
	communityRepo := repositories.NewCommunityRepository(db)
	communtiyService := services.NewCommunityService(communityRepo)
	communtiyHandler := handlers.NewCommunityHandler(communtiyService)

	//event member
	eventMemberRepo := repositories.NewEventMemberRepository(db)
	eventMemberService := services.NewEventMemberService(eventMemberRepo)
	eventMemberHandler := handlers.NewEventMemberHandler(eventMemberService)

	EventRoutes(router, eventHandler, eventMemberHandler, tokenRepo)

	// Route profile user.
	UserRoutes(router, userHandler, tokenRepo)

	//route
	AuthRoutes(router, authHandler)

	CommunityRoutes(router, communtiyHandler, tokenRepo)

	return router
}
