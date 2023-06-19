package user

import (
	"os"
	"time"

	"github.com/SantiiRepair/biosurf-api/db"
	"github.com/dgrijalva/jwt-go"
	fiber "github.com/gofiber/fiber/v2"
)

func getTokenInfo(token *jwt.Token) jwt.MapClaims {
	claims := token.Claims.(jwt.MapClaims)
	return claims
}

func HandleGoogle(c *fiber.Ctx) error {
	var users User
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

	claims := getTokenInfo(token)
	if claims == nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not verify token",
			"error":   err,
		})
	}

	loc, _ := time.LoadLocation("Europe/Madrid")
	date := time.Now().In(loc)

	if oauth2.Action == "login" {
		email := claims["email"].(string)
		sub := claims["sub"].(string)

		db.DB.Where("email = ?", email).First(&users)
		if len(users.GoogleAccount) == 0 {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "Google account not found",
			})
		}

		if users.GoogleID != sub {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "Illegal authentication",
			})
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString(secret)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Could not login",
			})
		}

		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "Login successful",
			"session": t,
		})
	}

	if oauth2.Action == "register" {
		email := claims["email"].(string)
		name := claims["name"].(string)
		family_name := claims["family_name"].(string)
		sub := claims["sub"].(string)

		db.DB.Where("email = ?", email).First(&users)
		if len(users.GoogleAccount) > 0 {
			c.Status(fiber.StatusConflict)
			return c.JSON(fiber.Map{
				"message": "Google account exists",
			})
		}

		users := &User{
			Name:      name,
			Lastname:  family_name,
			Email:     email,
			GoogleID:  sub,
			CreatedAt: date.String(),
			UpdatedAt: date.String(),
		}

		db.DB.Create(&users)
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "Register successful",
		})
	}

	c.Status(fiber.StatusForbidden)
	return c.JSON(fiber.Map{
		"message": date,
	})
}
