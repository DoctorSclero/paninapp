package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

// User registration

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (u *User) Create(ctx *gin.Context) {
	var create_request CreateUserRequest
	err := ctx.BindJSON(&create_request)
}
