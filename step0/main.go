package step0

import (
	"github.com/gofiber/fiber/v2"
)

func Step0() {
	app := fiber.New()
	app.Listen(":8000")
}
