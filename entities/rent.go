package entities

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	UserID       uint      `json:"user_id" form:"user_id"`
	BookID       uint      `json:"book_id" form:"book_id"`
	ReturnDate   time.Time `json:"return_date" form:"return_date"`
	ReturnStatus string    `json:"return_status" form:"return_status"`
	Address      Address   `gorm:"foreignKey:ID;references:ID" json:"address" form:"address"`
	User         User      `gorm:"foreignKey:UserID;references:ID"`
	Book         Book      `gorm:"foreignKey:BookID;references:ID"`
}
