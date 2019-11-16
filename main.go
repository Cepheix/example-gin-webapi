package main

import (
	"github.com/Cepheix/example-gin-webapi/common"
	"github.com/Cepheix/example-gin-webapi/products"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	db := common.Init()

	defer db.Close()

	ApplyMigrations(db)

	engine := gin.Default()

	RegisterModules(engine)

	engine.Run()
}

func RegisterModules(engine *gin.Engine) {
	api := engine.Group("/api")

	products.RegisterProductsModule(api.Group("/products"))
}

func ApplyMigrations(db *gorm.DB) {
	products.Migrate(db)
}
