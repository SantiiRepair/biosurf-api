package user

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"google.golang.org/api/oauth2/v2"
)

func HandleGoogle(c *fiber.Ctx) error {
	idToken := c.FormValue("googleToken")
	oauth2Service, err := oauth2.New(&http.Client{})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not auth",
		})
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	decode, err := tokenInfoCall.Do()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not auth",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": decode,
	})
}
