package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber_workshop/step3/controllers"
	"gorm.io/gorm"
	"fmt"
)

type Users struct {
	gorm.Model
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Router(app *fiber.App, db *gorm.DB) {
	app.Get("/home", controllers.Home)

	app.Post("/register", func(c *fiber.Ctx) error {
		tmp_user := &Users{}

		if err := c.BodyParser(tmp_user); err != nil {
			return err
		}

		db.AutoMigrate(&Users{})
		var result bool

		db.Raw("SELECT true FROM Users WHERE (username = ? AND password = ?) OR (username = ?)",
			tmp_user.Username, tmp_user.Password, tmp_user.Username).Scan(&result)

		if result == true {
			fmt.Println("HERE")
			return c.Status(401).SendString("Invalid username or password")
		} else {
			fmt.Println("HERE2")
			db.Create(&Users{Username: tmp_user.Username, Password: tmp_user.Password})
			return c.Status(200).SendString("Account successfully created")
		}
	})
}


