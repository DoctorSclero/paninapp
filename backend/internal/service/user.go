package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pietroballarin.com/paninup-backend/internal/model"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// User registration

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (u *UserService) Create(ctx *gin.Context) {
	var create_request CreateUserRequest
	// Checking request format
	if err := ctx.ShouldBindJSON(&create_request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Creating user
	user, err := model.NewUser(create_request.Email, create_request.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// Saving user in database
	u.db.Create(&user)
	ctx.JSON(200, gin.H{"user": user})
}
