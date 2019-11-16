package products

import "github.com/gin-gonic/gin"

func RegisterProductsModule(router *gin.RouterGroup) {
	router.POST("", CreateProduct)
}
