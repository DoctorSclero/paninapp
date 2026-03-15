package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pietroballarin.com/paninup-backend/internal/auth"
	"pietroballarin.com/paninup-backend/internal/types"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

// User registration

type RegistrationRequest struct {
	Email    string         `json:"email" binding:"required,email"`
	Password string         `json:"password" binding:"required,min=8"`
	Role     types.UserRole `json:"role"`
}

func (u *Service) Register(ctx *gin.Context) {
	var create_request RegistrationRequest
	// Checking request format
	if err := ctx.ShouldBindJSON(&create_request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Creating user
	user, err := New(create_request.Email, create_request.Password, create_request.Role)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// Saving user in database
	if err := u.db.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Could not create user"})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(200, gin.H{"user": user, "token": token})
}

// User login

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (u *Service) Login(ctx *gin.Context) {

	var login_request LoginRequest
	// Checking request format
	if err := ctx.ShouldBindJSON(&login_request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Finding user in database
	var user User
	u.db.Where("email = ?", login_request.Email).First(&user)

	// Checking if user exists and password is correct
	if user.ID == 0 || !user.CheckPassword(login_request.Password) {
		ctx.JSON(401, gin.H{"error": "User not found or incorrect password"})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	// Sending confirmation response
	ctx.JSON(200, gin.H{"user": user.Email, "role": user.Role, "token": token})
}
