package controllers

import (
	"detail-products/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// DetailProductHandler maneja la obtención de detalles de un producto
func DetailProductHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// Obtener el producto desde el modelo
		product, err := models.GetProductByID(db, id)
		if err != nil {
			http.Error(w, "Error al obtener el producto: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Si no se encuentra el producto
		if product == nil {
			http.Error(w, "Producto no encontrado", http.StatusNotFound)
			return
		}

		// Configurar los encabezados de la respuesta
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Codificar la respuesta en JSON
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, "Error al procesar la respuesta", http.StatusInternalServerError)
		}
	}
}
