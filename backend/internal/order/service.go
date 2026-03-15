package order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pietroballarin.com/paninup-backend/internal/sandwich"
	"pietroballarin.com/paninup-backend/internal/types"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

type OrderItemRequest struct {
	SandwichID uint `json:"sandwich_id" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required,min=1"`
}

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items" binding:"required,min=1"`
}

func (s *Service) Create(ctx *gin.Context) {
	// Get UserID from middleware
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var req CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Start transaction
	err := s.db.Transaction(func(tx *gorm.DB) error {
		order := Order{
			UserID: userID.(uint),
			Status: StatusPending,
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		var total float64
		for _, itemReq := range req.Items {
			var sw sandwich.Sandwich
			if err := tx.First(&sw, itemReq.SandwichID).Error; err != nil {
				return err
			}

			item := OrderItem{
				OrderID:    order.ID,
				SandwichID: sw.ID,
				Quantity:   itemReq.Quantity,
				Price:      sw.Price,
			}

			if err := tx.Create(&item).Error; err != nil {
				return err
			}

			total += sw.Price * float64(itemReq.Quantity)
		}

		// Update order total
		if err := tx.Model(&order).Update("total", total).Error; err != nil {
			return err
		}

		ctx.JSON(200, gin.H{"order_id": order.ID, "total": total})
		return nil
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func (s *Service) List(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	userRole, _ := ctx.Get("role")
	role := userRole.(types.UserRole)

	var orders []Order
	query := s.db.Preload("Items.Sandwich")

	if role == types.RoleManager {
		// Manager: all uncompleted orders (Pending or Confirmed)
		query = query.Where("status IN ?", []OrderStatus{StatusPending, StatusConfirmed})
	} else {
		// Consumer: their own orders (Pending, Confirmed, or Completed)
		query = query.Where("user_id = ? AND status IN ?", userID, []OrderStatus{StatusPending, StatusConfirmed, StatusCompleted})
	}

	if err := query.Find(&orders).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Could not retrieve orders"})
		return
	}

	ctx.JSON(200, gin.H{"orders": orders})
}
