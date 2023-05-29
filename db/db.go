package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "biosurf_adm:biosurf1234@tcp(db4free.net:3306)/biosurf_test")
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %s", err.Error())
	}

	return db, nil
}
