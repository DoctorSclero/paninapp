package sandwich

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

type CreateSandwichRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

func (s *Service) Create(ctx *gin.Context) {
	var req CreateSandwichRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sw := Sandwich{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	if err := s.db.Create(&sw).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Could not create sandwich"})
		return
	}

	ctx.JSON(200, gin.H{"sandwich": sw})
}

func (s *Service) List(ctx *gin.Context) {
	var sandwiches []Sandwich
	if err := s.db.Find(&sandwiches).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Could not retrieve sandwiches"})
		return
	}

	ctx.JSON(200, gin.H{"sandwiches": sandwiches})
}
