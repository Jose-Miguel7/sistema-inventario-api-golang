package models

import (
	"github.com/Jose-Miguel7/api-pos-go/internal/database"
)

func CreateProduct(product *Product) error {
	db, _ := database.GetDatabase()
	db.Create(product)
	return nil
}

func GetProduct(id uint) *Product {
	db, _ := database.GetDatabase()
	var product Product
	db.First(&product, id)

	return &product
}
