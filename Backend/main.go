// main.go
package main

import (
	"log"
	"net/http"
	"p3db/db"
	"p3db/handlers"
	"p3db/repositories"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Configuración de la base de datos
	config := db.NewDBConfig()
	config.Host = "localhost"
	config.Port = "5435"
	config.User = "dbuser"
	config.Password = "dbpassword"
	config.DBName = "p3db"

	// Obtener instancia de la base de datos
	dbInstance, err := db.GetDBInstance(config)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer dbInstance.Close()

	// Inicializar repositorios
	userRepo := repositories.NewUserRepository(dbInstance)

	// Inicializar handlers
	userHandler := handlers.NewUserHandler(userRepo)

	// Configurar enrutador
	router := mux.NewRouter()

	// Configurar endpoints
	router.HandleFunc("/test/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/test/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/test/users/{id}", userHandler.GetUser).Methods("GET")

	// Iniciar servidor
	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}