package bootstrap

import (
	"github.com/ahmadhdr/go-fiber-tutorial/routes"
	"github.com/gofiber/fiber"
)

func Start() {
	app := fiber.New()
	routes.Setup(app)
	app.Listen(3000)
}
