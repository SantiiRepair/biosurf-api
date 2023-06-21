package user

import (
	"os"
	"strings"

	"github.com/SantiiRepair/biosurf-api/db"
	"github.com/dgrijalva/jwt-go"
	fiber "github.com/gofiber/fiber/v2"
)

func HandleInfo(c *fiber.Ctx) error {
	var users User
	authHeader := c.Get("Authorization")
	secret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	if authHeader == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Authorization token not found")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid authorization token")
	}

	token := authParts[1]
	fun, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "Invalid token server",
				"token":   token,
			}), nil
		}

		return secret, nil
	})

	claims := fun.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	if id == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Empty token data")
	}

	db.DB.Where("id = ?", id).First(&users)
	if users.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Account not found",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"name":             users.Name,
		"lastname":         users.Lastname,
		"email":            users.Email,
		"google_account":   users.GoogleAccount,
		"facebook_account": users.FacebookAccount,
		"ipv4":             users.Ipv4,
	})
}
