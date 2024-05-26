package models

import (
	"github.com/Jose-Miguel7/api-pos-go/internal/database"
)

func CreateProduct(product *Product) error {
	db, _ := database.GetDatabase()
	db.Create(product)
	return nil
}

func GetProduct(id uint) (*Product, error) {
	db, _ := database.GetDatabase()
	var product Product
	result := db.First(&product, id)

	if result.Error != nil {
		return &product, result.Error
	}

	return &product, nil
}
