package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/db"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/handlers"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/repositories"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/routes"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/services"
	configs "github.com/gattini0928/Learning-Go-JWT-AUTH/internal/configs"

	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	conn := db.Connect()

	
	cfg := configs.LoadDBConfig()
	secret := []byte(cfg.JWTSecret)

	userRepo := repositories.NewUserRepository(conn)
	userService := services.NewUserService(userRepo, secret)
	userHandler := handlers.NewUserHandler(userService) 

	mux := http.NewServeMux()

	routes.RegisterUserRoute(mux, userHandler)
	routes.ProtectedRoute(mux, secret)

	port := os.Getenv("API_PORT")

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}