package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber_workshop/step3/controllers"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Router(app *fiber.App, db *gorm.DB) {
	app.Get("/home", controllers.Home)

	app.Post("/login", func(c *fiber.Ctx) error {
		tmp_user := &Users{}

		if err := c.BodyParser(tmp_user); err != nil {
			return err
		}

		db.AutoMigrate(&Users{})
		var result bool

		db.Raw("SELECT true FROM Users WHERE username = ? AND password = ?",
			tmp_user.Username, tmp_user.Password).Scan(&result)

		if result == false {
			return c.Status(401).SendString("Invalid username or password")
		} else {
			return c.Status(200).SendString("Logged in successfully")
		}
	})
}


