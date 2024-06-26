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
	product, err := models.GetProduct(uint(idProduct))

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Println(err)
		var manageError ManageError
		json.Unmarshal([]byte(`{"error": "No se encontró el producto}`), &manageError)
		json.NewEncoder(w).Encode(&manageError)
	} else {
		json.NewEncoder(w).Encode(product)
	}
}
