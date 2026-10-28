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

	router.Use(middleware.Cors)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Static("/uploads", "./uploads")

	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)
	passwordRepo := repositories.NewPasswordRepository(db)

	authService := services.NewAuthService(userRepo)
	passwordService := services.NewPasswordService(userRepo, passwordRepo)

	authHandler := handlers.NewAuthHandler(authService, tokenRepo, passwordService)

	userHandler := handlers.NewUserHandler(*userRepo)

	eventRepo := repositories.NewEventRepository(db, redisClient)
	eventService := services.NewEventService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)

	eventMemberRepo := repositories.NewEventMemberRepository(db)
	eventMemberService := services.NewEventMemberService(eventMemberRepo)
	eventMemberHandler := handlers.NewEventMemberHandler(eventMemberService)

	communityRepo := repositories.NewCommunityRepository(db)
	communityMemberRepo := repositories.NewCommunityMemberRepository(db)

	communityService := services.NewCommunityService(communityRepo)
	communityMemberService := services.NewCommunityMemberService(communityMemberRepo)

	communityHandler := handlers.NewCommunityHandler(communityService)
	communityMemberHandler := handlers.NewCommunityMemberHandler(communityMemberService, communityMemberRepo)

	notificationRepo := repositories.NewNotificationRepository(db)
	notificationService := services.NewNotificationService(notificationRepo)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	EventRoutes(router, eventHandler, eventMemberHandler, tokenRepo)
	UserRoutes(router, userHandler, tokenRepo)
	AuthRoutes(router, authHandler)
	CommunityRoutes(router, communityHandler, communityMemberHandler, tokenRepo)
	NotificationRoutes(router, notificationHandler, tokenRepo)

	return router
}
