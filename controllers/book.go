package controllers

import (
	"github.com/ahmadhdr/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")

	var book Book
	db.Find(&book, "id = ?", id)
	if book.Title == "" {
		c.Send("Book not found with given id")
	}
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
	}
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var book Book

	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No Book found with given id")
	}
	db.Delete(&book)
	c.Send("Book succesfuly deleted")

}

func UpdateBook(c *fiber.Ctx) {
	c.Send("Update Book")
}
