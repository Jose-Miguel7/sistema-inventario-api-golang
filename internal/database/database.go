package database

import (
	"github.com/Jose-Miguel7/api-pos-go/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	_, err := gorm.Open(sqlite.Open(config.Settings.DbPath), &gorm.Config{})
	if err != nil {
		panic("database doesn't work")
	}
}

func GetDatabase() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(config.Settings.DbPath), &gorm.Config{})
}
