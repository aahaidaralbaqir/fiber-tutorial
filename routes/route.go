package routes

import (
	"github.com/ahmadhdr/go-fiber-tutorial/controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/book", controllers.GetBooks)
	app.Get("/api/v1/book/:id", controllers.GetBook)
	app.Post("/api/v1/book", controllers.NewBook)
	app.Delete("/api/v1/book/:id", controllers.DeleteBook)
}
