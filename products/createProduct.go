package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(context *gin.Context) {
	var createProductRequest CreateProductRequest

	if err := context.ShouldBindJSON(&createProductRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productEntity := ProductEntity{}
	productEntity.Name = createProductRequest.Name
	productEntity.Price = createProductRequest.Price

	if err := AddProduct(&productEntity); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, productEntity.ID)
}

type CreateProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}
