package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jose-Miguel7/api-pos-go/internal/config"
	"github.com/Jose-Miguel7/api-pos-go/internal/handlers"
)

func InitServer() {
	route := http.NewServeMux()

	handlers.GetHandlers(route)

	server := &http.Server{
		Addr:    config.Settings.Port,
		Handler: route,
	}
	fmt.Printf("Running server in port %s\n", config.Settings.Port)
	log.Fatal(server.ListenAndServe())
}
