package verificetion

import (
	"fmt"
	"learn/database"
	"learn/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Verification(c *fiber.Ctx) error {
	tokenString := c.Request().Header.Peek("Authorization")
	token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("Secret Pass hhhh"), nil
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exo"].(float64) {
			return c.Status(400).JSON(http.StatusUnauthorized)
		}
		var user models.User
		database.Database.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			return c.Status(400).JSON("undefind user")
		}
		//c.Set("user", string(user))
		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		return c.Status(200).JSON("err i dont know ????")
	}
	return nil
}
