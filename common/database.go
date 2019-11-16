package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "my-database.db")
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}
