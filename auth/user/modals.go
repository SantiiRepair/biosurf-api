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

	statement, err := db.Prepare("INSERT INTO users (name, lastname, email, password) VALUES (?,?,?,?)")
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.LastName, user.Email, user.Password)
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
	query := db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if query != nil {
		if query == sql.ErrNoRows {
			return nil, fmt.Errorf("User with email %s not found", email)
		}
		return nil, query
	}
	return &user, nil
}
