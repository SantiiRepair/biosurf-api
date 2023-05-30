package user

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterData struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
