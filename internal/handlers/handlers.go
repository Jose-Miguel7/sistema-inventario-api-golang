package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Jose-Miguel7/api-pos-go/internal/models"
)

func POSTCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	err := models.CreateProduct(&product)
	if err != nil {
		log.Println("Hubo un error al crear el producto")
	}

	fmt.Fprintf(w, "Se creo el producto %s", product.Name)
}

func GETProduct(w http.ResponseWriter, r *http.Request) {
	idProduct, _ := strconv.Atoi(r.PathValue("id"))
	product := models.GetProduct(uint(idProduct))

	fmt.Fprintf(w, "El producto %s se creo el %s", product.Name, product.CreatedAt)
}

func GetHandlers(route *http.ServeMux) {
	route.HandleFunc("POST /api/product/", POSTCreateProduct)
	route.HandleFunc("GET /api/product/{id}/", GETProduct)
}
