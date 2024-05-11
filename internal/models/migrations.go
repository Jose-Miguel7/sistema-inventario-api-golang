package models

import "github.com/Jose-Miguel7/api-pos-go/internal/database"

func InitMigrations() {
	db, _ := database.GetDatabase()
	db.AutoMigrate(&Product{}, &User{}, &Order{})
}
