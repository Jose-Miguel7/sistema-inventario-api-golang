package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Total    string
	Quantity uint
	Product  uint `gorm:"index"`
}
