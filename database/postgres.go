package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(host, port, user, pass, name string) (*gorm.DB, error) {
	dsn := "host=" + host + " port=" + port + " user=" + user + " dbname=" + name + " password=" + pass + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
