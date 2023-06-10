package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() error{
	dsn := "biosurf_adm:biosurf1234@tcp(db4free.net:3306)/biosurf_test"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	DB = db
	
	return err
}
