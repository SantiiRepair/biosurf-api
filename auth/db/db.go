package db

import (
    "database/sql"
)

db, err := sql.Open("mysql", "usuario:contraseña@tcp(direccion:puerto)/nombre_de_la_base_de_datos")
if err != nil {
        return err
    }
    defer db.Close()