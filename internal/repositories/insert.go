package repositories

import (
	"database/sql"
	"github.com/gattini0928/Learning-Go-JWT-AUTH/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB:db}
}

func (r *UserRepository) InsertUser(user models.User) error {
	_, err := r.DB.Exec(`
		INSERT INTO users (name, email, password) 
		VALUES ($1,$2,$3)`, 
		user.Name,
		user.Email,
		user.Password,
	)
	return err
}
