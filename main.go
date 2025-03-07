package main

import (
	"detail-products/controllers"
	"detail-products/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Inicializar la base de datos
	db, err := models.InitDB()
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	// Verificar la conexión a la base de datos
	if err := db.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	// Crear un enrutador de mux
	router := mux.NewRouter()

	// Registrar el endpoint para obtener detalles del producto
	router.HandleFunc("/products/{id}", controllers.DetailProductHandler(db)).Methods("GET")

	// Configuración explícita de CORS para permitir solicitudes de todos los orígenes
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                                 // Permitir solicitudes de cualquier origen
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           // Métodos permitidos
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"}, // Encabezados permitidos
		ExposedHeaders:   []string{"Content-Type"},                                      // Encabezados expuestos
		AllowCredentials: false,                                                         // No permitir credenciales (cookies, headers de autenticación)
	}).Handler(router)

	// Iniciar el servidor
	log.Println("Servidor iniciado en :3009")
	log.Fatal(http.ListenAndServe(":3009", handler))
}
