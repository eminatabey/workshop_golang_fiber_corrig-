package step1

import (
	"github.com/gofiber/fiber/v2"
	"fiber_workshop/step1/routes"
)

func Step1() {
	app := fiber.New()

	routes.Router(app)
	app.Listen(":8000")
}