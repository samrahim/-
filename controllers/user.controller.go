package controllers

import (
	"learn/database"
	"learn/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var body struct {
		Email    string
		Password string
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"failed to parse this data": err.Error()})
	}
	database.Database.DB.Create(&models.User{Email: body.Email, Password: string(hash)})
	return c.Status(200).JSON("user created")
}
func Login(c *fiber.Ctx) error {
	var body struct {
		Email    string
		Password string
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	database.Database.DB.Find(&user, "email=?", body.Email)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(200).JSON(err.Error())
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": &user.ID,
		"exo": time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("Secret Pass hhhh"))
	c.Response().Header.Set("Authorization", tokenString)
	return c.Status(200).JSON("")

}
