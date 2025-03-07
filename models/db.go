package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/denisenkom/go-mssqldb" // Driver de SQL Server
)

// InitDB inicializa la conexión a la base de datos
func InitDB() (*sqlx.DB, error) {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Obtener variables de entorno
	dbHost := os.Getenv("SQL_DB_HOST")
	dbUser := os.Getenv("SQL_DB_USER")
	dbPassword := os.Getenv("SQL_DB_PASSWORD")
	dbName := os.Getenv("SQL_DB_NAME")

	// Validar que las variables están presentes
	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Faltan variables de entorno en el archivo .env")
	}

	// Construir la cadena de conexión
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		dbUser, dbPassword, dbHost, dbName)

	// Conectar a SQL Server
	db, err := sqlx.Connect("sqlserver", connStr)
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	return db, nil
}
