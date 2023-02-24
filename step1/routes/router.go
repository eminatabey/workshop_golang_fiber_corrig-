package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber_workshop/step1/controllers"
)

func Router(app *fiber.App) {
	app.Get("/home", controllers.Home)
}
