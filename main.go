package main

import (
	"github.com/ahmadhdr/go-fiber-tutorial/application"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	application.Start()
}
