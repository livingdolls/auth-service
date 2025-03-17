package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/livingdolls/auth-service/internal/adapter/dto"
	"github.com/livingdolls/auth-service/internal/core/port"
)

type UserHandler struct {
	userService port.UserPortService
}

func NewUserHandler(userService port.UserPortService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) GetUserByEmail(c *gin.Context) {
	var req dto.UserGetByEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userService.GetDetailUserByEmail(req.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}
