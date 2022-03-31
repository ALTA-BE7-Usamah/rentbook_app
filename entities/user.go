package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"unique;not null" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Book     []Book `gorm:"foreignKey:UserID;references:ID"`
	Rent     []Rent `gorm:"foreignKey:UserID;references:ID"`
}
