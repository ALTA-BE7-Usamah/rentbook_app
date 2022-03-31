package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	UserID    uint   `json:"user_id" form:"user_id"`
	Title     string `json:"title" form:"title"`
	Catagory  string `json:"catagory" form:"catagory"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Status    string `gorm:"default:avalaible" json:"status" form:"status"`
	Rent      *Rent  `gorm:"foreignKey:BookID;references:ID"`
}
