package db

import (
    "database/sql"
    "fmt"
)

func Connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "user:password@tcp(address:port)/database_name")
    if err != nil {
        return nil, fmt.Errorf("Error connecting to database: %s", err.Error())
    }

    return db, nil
}