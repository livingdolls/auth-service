package router

import (
	"github.com/gin-gonic/gin"
	"github.com/livingdolls/auth-service/internal/adapter/handler"
)

func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	AuthRoutes(r, authHandler)

	return r
}
