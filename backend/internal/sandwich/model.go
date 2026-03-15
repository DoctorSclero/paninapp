package sandwich

import (
	"gorm.io/gorm"
)

type Sandwich struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
}
