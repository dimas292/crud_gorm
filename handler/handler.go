package handler

import (
	"latihan2/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// NewCreateProduct adalah handler untuk membuat produk baru
func (d *Database) NewCreateProduct(ctx *gin.Context) {
	var req entity.Product

	// Binding data dari request body
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"messgae": "ERR BAD REQUEST",
			"error":   "`field` tidak boleh kosong",
		})
		return
	}

	// Menyimpan data ke database
	if err := d.DB.Create(&req).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Mengirimkan response sukses
	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"messages": "CREATE SUCCESS",
		"payload":  req,
	})
}

func (d *Database) NewGetAll(ctx *gin.Context) {

	var product []entity.Product

	err := d.DB.Find(&product).Error

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "ERR NOT FOUND",
			"error":   "data tidak ditemukan",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "GET ALL SUCCESS",
		"payload": product,
	})
}

func (d *Database) NewGetOneByID(ctx *gin.Context) {

	var product entity.Product

	id := ctx.Param("id")

	err := d.DB.Where("id=?", id).First(&product).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ERR NOT FOUND",
			"error":   "data tidak ditemukan",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "GET DATA SUCCESS",
		"payload": product,
	})
}

func (d *Database) NewUpdate(ctx *gin.Context) {
	var product entity.Product
	var req entity.Product

	id := ctx.Param("id")

	if err := d.DB.Where("id=?", id).First(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   "`nama` tidak boleh kosong",
		})
	}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "gagal update",
		})
	}
	product.Name = req.Name
	product.Category = req.Category

	err := d.DB.Save(&product)

	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "UPDATE SUCCESS",
	})

}

func (d *Database) NewDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	var product []entity.Product

	if err := d.DB.Find(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   "`nama` tidak boleh kosong",
		})
	}

	err := d.DB.Delete(&product, id)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "DELETE DATA SUCCESS",
	})

}
