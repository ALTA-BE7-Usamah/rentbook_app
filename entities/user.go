package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Book     []Book `gorm:"foreignKey:UserID;references:ID"`
	Rent     []Rent `gorm:"foreignKey:UserID;references:ID"`
}
