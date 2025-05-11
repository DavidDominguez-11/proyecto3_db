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
	artistRepo := repositories.NewArtistRepository(dbInstance)

	// Inicializar handlers
	userHandler := handlers.NewUserHandler(userRepo)
	artistHandler := handlers.NewArtistHandler(artistRepo)

	// Configurar enrutador
	router := mux.NewRouter()

	// Configurar endpoints
	//User
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	//Artist
	router.HandleFunc("/artists", artistHandler.GetArtists).Methods("GET")
	router.HandleFunc("/artists", artistHandler.CreateArtist).Methods("POST")
	router.HandleFunc("/artists/{id}", artistHandler.GetArtist).Methods("GET")
	router.HandleFunc("/artists/{id}", artistHandler.UpdateArtist).Methods("PUT")
	router.HandleFunc("/artists/{id}", artistHandler.DeleteArtist).Methods("DELETE")


	// Iniciar servidor
	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}