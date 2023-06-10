package report

import (
	fiber "github.com/gofiber/fiber"
)

func HandleReport(c *fiber.Ctx) {
	text := c.FormValue("text")

	response := fiber.Map{
		"message": "The text does not contain obscene words",
		"obscene": false,
	}

	if isProfanity(text) {
		response["message"] = "The text contains an obscene word"
		response["obscene"] = true

		c.Status(fiber.StatusAccepted)
		c.JSON(fiber.Map{"response": response})
		return
	}
	c.Status(fiber.StatusAccepted)
	c.JSON(fiber.Map{"response": response})
}
