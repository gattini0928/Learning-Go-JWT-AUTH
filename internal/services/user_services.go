package services

import (
	"errors"

	config "github.com/gattini0928/Learning-Go-JWT-AUTH/internal/configs"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/models"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/repositories"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/services/auth"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenFailed = errors.New("token failed")
)

type UserService struct {
	Repo *repositories.UserRepository
	Secret []byte
}

func NewUserService(rp *repositories.UserRepository, secret []byte) *UserService {
	return &UserService{Repo:rp, Secret: secret}
}

func (u *UserService) CreateUser(user models.User) error {
	err := ValidateName(user.Name)
	if err != nil {
		return err
	}

	err = ValidateEmail(user.Email)
	if err != nil {
		return err
	}

	err = ValidatePassword(user.Password)
	if err != nil {
		return err
	}

	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return ErrInvalidPassword
	}

	user.Password = hashPassword

	return u.Repo.InsertUser(user)
}

func (s *UserService) Login(email, password string) (models.User, string, error) {

	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, "", ErrUserNotFound
	}

	if !CheckPasswordHash(password, user.Password) {
		return models.User{}, "", ErrInvalidPassword
	}

	cfg := config.LoadDBConfig()

	token, err := auth.CreateJWT(s.Secret, user.ID, cfg.JWTExpirationInSeconds)
	if err != nil {
		return models.User{}, "" , ErrTokenFailed
	}

	return user, token, nil
}