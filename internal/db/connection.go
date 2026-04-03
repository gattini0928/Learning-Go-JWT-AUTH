package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/configs"
)

func Connect() *sql.DB {
	cfg := config.LoadDBConfig()
	connStr := cfg.ConnectionString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}