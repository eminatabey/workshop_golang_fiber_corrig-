package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber_workshop/step5/controllers"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID int `json:"id"`
	usersname string `json:"usersname"`
	Password string `json:"password"`
}

type Todos struct {
	gorm.Model
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status bool `json:"status"`
}

func Router(app *fiber.App, db *gorm.DB) {
	app.Get("/home", controllers.Home)

	// LOGIN REGISTER

	app.Post("/login", func(c *fiber.Ctx) error {
		tmp_users := &Users{}

		if err := c.BodyParser(tmp_users); err != nil {
			return err
		}

		db.AutoMigrate(&Users{})
		var result bool

		db.Raw("SELECT true FROM users WHERE usersname = ? AND password = ?",
			tmp_users.usersname, tmp_users.Password).Scan(&result)

		if result == false {
			return c.Status(401).SendString("Invalid usersname or password")
		} else {
			return c.Status(200).SendString("Logged in successfully")
		}
	})


	app.Post("/register", func(c *fiber.Ctx) error {
		tmp_users := &Users{}

		if err := c.BodyParser(tmp_users); err != nil {
			return err
		}

		db.AutoMigrate(&Users{})
		var result bool

		db.Raw("SELECT true FROM users WHERE (usersname = ? AND password = ?) OR (usersname = ?)",
			tmp_users.usersname, tmp_users.Password, tmp_users.usersname).Scan(&result)

		if result == true {
			return c.Status(401).SendString("Invalid usersname or password")
		} else {
			db.Create(&Users{usersname: tmp_users.usersname, Password: tmp_users.Password})
			return c.Status(200).SendString("Account successfully created")
		}
	})

	// TODOS

	app.Get("/list-todo", func(c *fiber.Ctx) error {
		db.AutoMigrate(&Todos{})
		var todo_list []Todos
		db.Find(&todo_list)
		return c.JSON(todo_list)
	})

}


