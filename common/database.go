package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../my-database.db")

	if err != nil {
		fmt.Println("Cannot open database")
	}

	return db
}
