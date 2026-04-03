package main

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/db"
)

func main() {
	godotenv.Load()
	conn := db.Connect()
	defer conn.Close()

	query := `
		CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(150) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL
		);
	`
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tabela users criado com sucesso")

}

