package controllers

import (
	"github.com/gin-gonic/gin"
	"pietroballarin.com/paninup-backend/internal/models"
	"pietroballarin.com/paninup-backend/internal/repositories"
)

type User struct {
	Repo *repositories.User
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (u *User) Register(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{}
	if err := user.CreateFromEmailAndPassword(req.Email, req.Password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
