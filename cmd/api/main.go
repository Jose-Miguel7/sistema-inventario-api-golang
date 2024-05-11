package main

import (
	"github.com/Jose-Miguel7/api-pos-go/internal/database"
	"github.com/Jose-Miguel7/api-pos-go/internal/models"
	"github.com/Jose-Miguel7/api-pos-go/internal/server"
)

func main() {
	database.InitDatabase()
	models.InitMigrations()
	server.InitServer()
}
