package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Product representa la estructura de un producto
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
}

// GetProductByID obtiene los detalles de un producto por su ID
func GetProductByID(db *sqlx.DB, id int) (*Product, error) {
	var product Product
	query := `SELECT id, name, description, price, quantity, category, brand FROM Products WHERE id = @id`
	// Utiliza el parámetro con nombre para SQL Server
	err := db.Get(&product, query, sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Producto no encontrado
		}
		return nil, err // Error al ejecutar la consulta
	}
	return &product, nil
}
