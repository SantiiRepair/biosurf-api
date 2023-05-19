package auth

import (
	db "gitlab.com/biosurf/biosurf-api/auth/db"
)

func CreateUser(user *User) error {
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
	user.ID = int(userID)

	return nil
}
