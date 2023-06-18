package user

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	fiber "github.com/gofiber/fiber/v2"
)

func getTokenInfo(token *jwt.Token) jwt.MapClaims {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		sub := claims["sub"].(string)
		name := claims["name"].(string)
		iat := claims["iat"].(float64)

		fmt.Printf("sub: %s, name: %s, iat: %.0f\n", sub, name, iat)
	}
	return claims
}

func HandleGoogle(c *fiber.Ctx) error {
	var oauth2 GoogleData
	err := c.BodyParser(&oauth2)
	secret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Wrong params",
		})
	}

	token, err := jwt.Parse(oauth2.JWTDataUser, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message":      "Invalid access token",
				"access_token": oauth2.JWTDataUser,
			}), err
		}
		return secret, err
	})

	o := getTokenInfo(token)
	fmt.Println(o)
	if o == nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not verify token",
			"error":   err,
		})
	}

	//if oauth2.Action == "login" {
	//	fmt.Println(decode)
	//}

	//if oauth2.Action == "register" {
	//	fmt.Println(decode)
	//}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": o,
	})
}
