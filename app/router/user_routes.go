package router

import (
	"github.com/gin-gonic/gin"
	"github.com/livingdolls/auth-service/internal/adapter/handler"
)

func UserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	user := r.Group("/user")
	{
		user.GET("/email", userHandler.GetUserByEmail)
	}
}
