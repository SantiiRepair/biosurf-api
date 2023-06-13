package user

import (
	fiber "github.com/gofiber/fiber/v2"
)

func HandleLogout(c *fiber.Ctx) error {

	c.ClearCookie("smsuances_session")
	c.JSON(fiber.Map{
		"message": "success",
	})

	return nil
}
