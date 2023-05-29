package user

import (
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
	
	var user User
	w := db.QueryRow("SELECT id, email, password, created_at, updated_at FROM users WHERE email=?", email)
	error := w.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if error != nil {
		return nil, err
	}
	return &user, nil
}
