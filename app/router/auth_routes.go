package router

import (
	"github.com/gin-gonic/gin"
	"github.com/livingdolls/auth-service/internal/adapter/handler"
)

func AuthRoutes(r *gin.Engine, authHandler *handler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
	}
}
