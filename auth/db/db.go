package db

import (
    "database/sql"
    "fmt"
)

func Connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "usuario:contraseña@tcp(direccion:puerto)/nombre_de_la_base_de_datos")
    if err != nil {
        return nil, fmt.Errorf("Error conectando a la base de datos: %s", err.Error())
    }

    return db, nil
}