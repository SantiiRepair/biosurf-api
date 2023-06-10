package user

import (
	"time"

	fiber "github.com/gofiber/fiber"
)

func HandleLogout(c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.JSON(fiber.Map{
		"message": "success",
	})
}
