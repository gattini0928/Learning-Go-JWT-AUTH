package models

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}