package order

import (
	"gorm.io/gorm"
	"pietroballarin.com/paninup-backend/internal/sandwich"
	"pietroballarin.com/paninup-backend/internal/user"
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusConfirmed OrderStatus = "confirmed"
	StatusCompleted OrderStatus = "completed"
	StatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	gorm.Model
	UserID uint        `json:"user_id" gorm:"not null"`
	User   user.User   `json:"-" gorm:"foreignKey:UserID"`
	Items  []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Status OrderStatus `json:"status" gorm:"not null;default:'pending'"`
	Total  float64     `json:"total" gorm:"not null"`
}

type OrderItem struct {
	gorm.Model
	OrderID    uint              `json:"order_id" gorm:"not null"`
	SandwichID uint              `json:"sandwich_id" gorm:"not null"`
	Sandwich   sandwich.Sandwich `json:"sandwich" gorm:"foreignKey:SandwichID"`
	Quantity   int               `json:"quantity" gorm:"not null;default:1"`
	Price      float64           `json:"price" gorm:"not null"` // Captured at order time
}
