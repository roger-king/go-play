package handlers

import (
	"encoding/json"
	"github.com/roger-king/go-ecommerce/pkg/models"
	"github.com/roger-king/go-ecommerce/pkg/utilities"
	"net/http"
)

func FindProductsController(w http.ResponseWriter, req *http.Request) {
	products, err := models.AllProducts()

	if err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "Error getting products")
	}
	// TODO: Handle Errors on Service side.
	utilities.RespondWithJSON(w, http.StatusOK, products)
	return
}

func CreateProductsController(w http.ResponseWriter, req *http.Request) {
	var p models.Product

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&p); err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer req.Body.Close()

	models.CreateProduct(p)

	utilities.RespondWithJSON(w, http.StatusCreated, p)
}
