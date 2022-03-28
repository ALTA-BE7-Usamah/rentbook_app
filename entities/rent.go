package entities

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	UserID     uint      `json:"user_id" form:"user_id"`
	BookID     uint      `json:"book_id" form:"book_id"`
	ReturnDate time.Time `json:"return_date" form:"return_date"`
	Address    Address   `gorm:"foreignKey:ID;references:ID" json:"address" form:"address"`
}
