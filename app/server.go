package app

import (
	"github.com/livingdolls/auth-service/app/router"
	"github.com/livingdolls/auth-service/internal/adapter/handler"
	"github.com/livingdolls/auth-service/internal/adapter/repository"
	"github.com/livingdolls/auth-service/internal/core/service"
	"github.com/livingdolls/auth-service/internal/infrastructure/db"
)

func StartServer() {

	database := db.ConnectDB()

	userRepo := repository.NewUserRepository(database)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	r := router.SetupRouter(authHandler)

	r.Run(":8080")
}
