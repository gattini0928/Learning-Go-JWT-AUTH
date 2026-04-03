package routes

import (
	"net/http"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/handlers"
)

func RegisterUserRoute(mux *http.ServeMux, h *handlers.UserHandler) {
	mux.HandleFunc("POST /signup", h.HandleSignup)
	mux.HandleFunc("POST /login", h.HandleLogin)
}