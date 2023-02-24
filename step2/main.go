package step2

import (
	"github.com/gofiber/fiber/v2"
	"fiber_workshop/step2/routes"
	"fiber_workshop/step2/config"
)

func Step2() {
	app := fiber.New()

	db := config.Connect_database()

	_ = db
	routes.Router(app)
	app.Listen(":8000")
}