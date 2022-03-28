package entities

type Address struct {
	ID     uint   `gorm:"primaryKey"`
	Street string `json:"street" form:"street"`
	City   string `json:"city" form:"city"`
	State  string `json:"state" form:"state"`
	Zip    string `json:"zip" form:"zip"`
}
