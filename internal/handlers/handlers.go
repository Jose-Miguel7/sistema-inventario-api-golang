package handlers

import (
	"net/http"
)

type ManageError struct {
	Error string
}

func GetHandlers(route *http.ServeMux) {
	route.HandleFunc("POST /api/product/", POSTCreateProduct)
	route.HandleFunc("GET /api/product/{id}/", GETProduct)
}
