package main

import (
	"latihan2/database"
	"latihan2/entity"
	"latihan2/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func main() {
	// config
	cfg, err := loadConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	// connect database
	db, err := database.ConnectPostgres(
		"localhost",
		"5432",
		"postgres",
		"root",
		"postgres",
	)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db not connected")
	}
	log.Println("db connected")
	// migrasi
	db.AutoMigrate(entity.Product{})
	// routing
	r := gin.New()
	router.SetUp(r, db)
	r.Run(cfg.App.Port)
}

func loadConfig(filename string) (conf entity.Config, err error) {
	f, err := os.ReadFile(filename)

	if err != nil {
		return
	}
	err = yaml.Unmarshal(f, &conf)
	return
}
