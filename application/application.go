package application

import (
	"github.com/ahmadhdr/go-fiber-tutorial/database"
	"github.com/ahmadhdr/go-fiber-tutorial/routes"
	"github.com/gofiber/fiber"
)

func Start() {
	app := fiber.New()
	routes.Setup(app)
	database.Setup()
	app.Listen(3000)
}
