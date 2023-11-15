package main

import (
	"learn/controllers"
	"learn/database"
	"learn/verificetion"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.GetDb()
	app := fiber.New()
	app.Post("/signup", controllers.SignUp)
	app.Post("/login", controllers.Login)
	app.Get("/users", verificetion.Verification, func(c *fiber.Ctx) error {
		return c.Status(200).JSON("helllo")
	})
	app.Listen(":8000")
}
