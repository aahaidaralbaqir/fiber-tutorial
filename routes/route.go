package routes

import (
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/book", func(c *fiber.Ctx) {
		c.Send("Hallo froom book again")
	})
}
