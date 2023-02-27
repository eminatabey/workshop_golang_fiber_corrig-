package step3

import (
	"github.com/gofiber/fiber/v2"
	"fiber_workshop/step3/routes"
	"fiber_workshop/step3/config"
)

func Step3() {
	app := fiber.New()

	db := config.Connect_database()

	_ = db

	routes.Router(app, db)
	app.Listen(":8000")
}