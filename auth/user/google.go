package user

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"google.golang.org/api/oauth2/v2"
)

func HandleGoogle(c *fiber.Ctx) (*oauth2.Tokeninfo, error) {
	idToken := c.FormValue("googleToken")
	oauth2Service, err := oauth2.New(&http.Client{})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil, c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	decode, err := tokenInfoCall.Do()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil, c.JSON(fiber.Map{
			"message": "Could not a",
		})
	}

	c.Status(fiber.StatusOK)
	return nil, c.JSON(fiber.Map{
		"message": decode,
	})
}
