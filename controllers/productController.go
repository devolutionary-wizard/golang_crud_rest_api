package controllers

import (
	"encoding/json"
	"golang-rest-api/database"
	"golang-rest-api/entities"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}
