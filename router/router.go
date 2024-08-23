package router

import (
	"latihan2/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetUp(router *gin.Engine, db *gorm.DB){

	dbHandler := handler.NewDatabaseHandler(db)
	
	router.POST("/products", dbHandler.NewCreateProduct)
	router.GET("/products", dbHandler.NewGetAll)
	router.GET("/product/:id", dbHandler.NewGetOneByID)
	router.PUT("/product/:id", dbHandler.NewUpdate)
	router.DELETE("/product/:id", dbHandler.NewDelete)

}