package main

import (
	"github.com/Cepheix/example-gin-webapi/common"
	"github.com/Cepheix/example-gin-webapi/products"
	"github.com/jinzhu/gorm"
)

func main() {
	db := common.Init()

	defer db.Close()

	ApplyMigrations(db)
}

func ApplyMigrations(db *gorm.DB) {
	products.Migrate(db)
}
