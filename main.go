package main

import (
	"learn/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	setAppRoutes(app)
	app.Listen(":8000")
}

func setAppRoutes(app *fiber.App) {
	app.Get("/products", routes.FindAll)
	app.Post("/product", routes.CreateProduct)
	app.Get("/products/:status", routes.GetByStatus)
	app.Get("/products/:min/:max", routes.FindAllByPrices)
}
