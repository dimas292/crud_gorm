package main

import (
	"latihan2/database"
	"latihan2/entity"
	"latihan2/handler"
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

	router := gin.New()

	dbHandler := &handler.Database{DB: db}

	router.POST("/products", dbHandler.NewCreateProduct)
	router.GET("/products", dbHandler.NewGetAll)
	router.GET("/product/:id", dbHandler.NewGetOneByID)
	router.PUT("/product/:id", dbHandler.NewUpdate)
	router.DELETE("/product/:id", dbHandler.NewDelete)

	router.Run(cfg.App.Port)
}

func loadConfig(filename string) (conf entity.Config, err error) {

	f, err := os.ReadFile(filename)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(f, &conf)
	return
}
