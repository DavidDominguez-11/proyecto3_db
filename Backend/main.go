package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5435 user=dbuser password=dbpassword dbname=p3db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error abriendo conexión:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	fmt.Println("Conexión exitosa a PostgreSQL!")
}
