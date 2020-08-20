package controllers

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("Single book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("New Bok")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete book")
}

func UpdateBook(c *fiber.Ctx) {
	c.Send("Update Book")
}
