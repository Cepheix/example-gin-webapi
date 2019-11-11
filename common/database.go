package common

import (
	"github.com/jinzhu/gorm",
	"fmt"
)

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../my-database.db")

	if err != nil {
		fmt.Println("Cannot open database")
	}

	return db;
}
