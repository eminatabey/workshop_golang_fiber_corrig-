package step6

import (
	"github.com/gofiber/fiber/v2"
	"fiber_workshop/step6/routes"
	"fiber_workshop/step6/config"
)

func Step6() {
	app := fiber.New()

	db := config.Connect_database()

	_ = db

	routes.Router(app, db)
	app.Listen(":8000")
}