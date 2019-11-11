package products

import (
	"github.com/jinzhu/gorm"
)

type ProductEntity struct {
	gorm.Model
	Name  string
	Price uint
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&ProductEntity{})
}
