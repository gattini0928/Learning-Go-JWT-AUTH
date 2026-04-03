package routes

import (
	"net/http"

	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/handlers"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/services/auth"
)

func ProtectedRoute(mux *http.ServeMux, secret []byte) {
	mux.Handle("GET /protected",
		auth.JWTMiddleware(secret, http.HandlerFunc(handlers.HandleProtected)),
	)
}