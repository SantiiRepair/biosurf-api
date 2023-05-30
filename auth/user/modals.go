package user

import (
	"database/sql"
	"fmt"

	db "github.com/SantiiRepair/biosurf-api/db"
)

func CreateUser(user *User) error {
	db, err := db.Connect()
	statement, err := db.Prepare("INSERT INTO usuarios (email, password) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int64(userID)

	return nil
}

func GetUserByEmail(email string) (*User, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}
	var user User
	w := db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if w != nil {
		if w == sql.ErrNoRows {
			return nil, fmt.Errorf("User with email %s not found", email)
		}
		return nil, w
	}
	return &user, nil
}
