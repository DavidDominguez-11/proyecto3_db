package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

// DBConfig contiene la configuración para la conexión a la base de datos
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Database representa la conexión a la base de datos y proporciona métodos de acceso
type Database struct {
	db *sql.DB
}

var (
	instance *Database
	once     sync.Once
)

// NewDBConfig crea una configuración con valores por defecto
func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:    "localhost",
		Port:    "5432", // Puerto por defecto de PostgreSQL
		SSLMode: "disable",
	}
}

// GetConnectionString genera el string de conexión
func (c *DBConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

// GetDBInstance devuelve una instancia singleton de Database
func GetDBInstance(config *DBConfig) (*Database, error) {
	var initErr error
	once.Do(func() {
		db, err := sql.Open("postgres", config.GetConnectionString())
		if err != nil {
			initErr = fmt.Errorf("error abriendo conexión: %v", err)
			return
		}

		// Verificamos que la conexión esté activa
		if err = db.Ping(); err != nil {
			initErr = fmt.Errorf("no se pudo conectar a la base de datos: %v", err)
			return
		}

		log.Println("Conexión exitosa a PostgreSQL!")
		instance = &Database{db: db}
	})

	if initErr != nil {
		return nil, initErr
	}

	return instance, nil
}

// Close cierra la conexión a la base de datos
func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

// GetDB devuelve la instancia subyacente de *sql.DB
func (d *Database) GetDB() *sql.DB {
	return d.db
}

// HealthCheck verifica si la base de datos está respondiendo
func (d *Database) HealthCheck() error {
	return d.db.Ping()
}