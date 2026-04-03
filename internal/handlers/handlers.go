package handlers

import (
	"errors"
	"net/http"
	"fmt"

	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/models"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/services"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/services/auth"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/utils"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler{
	return &UserHandler{Service:s}
}

func (h *UserHandler) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadGateway, err)
		return
	}
	err := h.Service.CreateUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "user created",
	})
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var input models.User
	utils.ParseJSON(r, &input)

	user, token, err := h.Service.Login(input.Email, input.Password)
	if err != nil {

		if errors.Is(err, services.ErrInvalidPassword) ||
			errors.Is(err, services.ErrUserNotFound) {

			utils.WriteError(w, http.StatusUnauthorized, errors.New("email ou senha inválidos"))
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := models.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}

	utils.WriteJSON(w, http.StatusOK, res)
}

func HandleProtected(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(auth.UserIDKey).(string)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	res := map[string]any{
		"message": "rota protegida funcionando",
		"userID":  userID,
	}

	utils.WriteJSON(w, http.StatusOK, res)
}