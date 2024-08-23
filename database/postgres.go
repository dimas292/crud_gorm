package database

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)



func ConnectPostgres(host, port, user, pass, dbname string)(db *gorm.DB, err error){

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, pass, dbname)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}

	return db, nil
}