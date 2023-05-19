package auth

import (
	"database/sql"

	db "github.com/SantiiRepair/biosurf-api/auth/db"
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
	user.ID = int(userID)

	return nil
}

func GetUserByEmail(email string) (*User, error) {
    db, err := db.Connect()
    row := db.QueryRow("SELECT id, email, password FROM usuarios WHERE email = ?", email)

    user := User{}
    err = row.Scan(&user.ID, &user.Email, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrUserNotFound
        }
        return nil, err
    }

    return &user, nil
}