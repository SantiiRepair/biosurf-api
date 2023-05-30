package user

import (
	"database/sql"
	"fmt"
	"strings"

	db "github.com/SantiiRepair/biosurf-api/db"
)

func CreateUser(user *User) error {
	db, error := db.Connect()
	if error != nil {
		return error
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (name, lastname, email, password, created_at, updated_at) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return fmt.Errorf("Email already in use")
		}
		return fmt.Errorf("Failed to execute statement: %v", err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Failed to get last insert ID: %v", err)
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {
	db, error := db.Connect()
	if error != nil {
		panic(error)
	}
	defer db.Close()
	var user User
	query := db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Password)
	if query == sql.ErrNoRows {
		return nil, fmt.Errorf("User with email %s not found", email)
	}

	return &user, nil
}
