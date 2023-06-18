package user

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	fiber "github.com/gofiber/fiber/v2"
)

func getTokenInfo(accessToken string) (*oauth2.Tokeninfo, error) {
	ctx := context.Background()
	file, err := os.Open("./config/credentials.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		return nil, err
	}

	var creds struct {
		Web struct {
			ClientID     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
			RedirectURL  string `json:"redirect_uris"`
		} `json:"web"`
	}

	buf := json.Unmarshal(buffer[:n], &creds)
	if buf != nil {
		return nil, buf
	}

	credsJSON, err := json.Marshal(creds)
	if err != nil {
		return nil, err
	}

	oauth2Service, err := oauth2.NewService(ctx, option.WithCredentialsJSON(credsJSON))
	if err != nil {
		return nil, err
	}

	tokenInfo, err := oauth2Service.Tokeninfo().AccessToken(accessToken).Do()
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}

func HandleGoogle(c *fiber.Ctx) error {
	var oauth2 GoogleData
	err := c.BodyParser(&oauth2)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Wrong params",
		})
	}

	if len(oauth2.AccessToken) == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":      "Invalid access token",
			"access_token": oauth2.AccessToken,
		})
	}

	o, err := getTokenInfo(oauth2.AccessToken)
	fmt.Println(o, err)
	if err != nil {
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
