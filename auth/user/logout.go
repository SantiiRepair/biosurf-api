package user

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

func HandleLogout(c *fiber.Ctx)error {
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
	
	return nil
}
