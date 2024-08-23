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
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Pass,
		cfg.DB.Name,
	)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db not connected")
	}

	log.Println("db connected")

	db.AutoMigrate(entity.Product{})

	r := gin.New()

	router.SetUp(r)

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
