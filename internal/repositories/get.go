package repositories

import "github.com/gattini0928/Learning-Go-JWT-AUTH/internal/models"

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	row := r.DB.QueryRow(`
		SELECT id, name, email, password 
		FROM users 
		WHERE email = $1`, email)

	var user models.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
