package routes

import (
	"learn/config"
	"learn/entities"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindAll(c *fiber.Ctx) error {
	prods := []entities.Product{}
	db, err := config.GetDb()
	if err != nil {
		return c.Status(400).JSON("failed to connect to db")
	}
	db.Find(&prods)
	return c.Status(200).JSON(prods)
}

func CreateProduct(c *fiber.Ctx) error {
	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON("failed to parse this data")
	}
	db, err := config.GetDb()
	if err != nil {
		return c.Status(400).JSON("can't connect to db")
	}
	db.Create(&product)
	return c.Status(200).JSON(fiber.Map{
		"product created": &product,
	})
}

func GetByStatus(c *fiber.Ctx) error {
	status := c.Params("status")
	b, err := strconv.ParseBool(status)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	prods := []entities.Product{}
	db, err := config.GetDb()
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	db.Find(&prods, "status = ?", b)

	return c.Status(200).JSON(&prods)
}

func FindAllByPrices(c *fiber.Ctx) error {
	min, err := c.ParamsInt("min")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	max, err2 := c.ParamsInt("max")
	if err2 != nil {
		return c.Status(400).JSON(err.Error())
	}
	prods := []entities.Product{}
	db, err3 := config.GetDb()
	if err3 != nil {
		return c.Status(400).JSON("failed to connect db")
	}
	db.Find(&prods, "price >= ? AND price <= ?", min, max)
	return c.Status(200).JSON(prods)
}
