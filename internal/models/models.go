package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

type Product struct {
	gorm.Model
	Name  string
	Price float32
}

type Order struct {
	gorm.Model
	Total    string
	Quantity uint
	Product  uint `gorm:"index"`
}
