package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	DBConn *gorm.DB
)

func Setup() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Database connection successfull opened")
}
