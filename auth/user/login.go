package user

import (
	"strconv"
	"time"

	db "github.com/SantiiRepair/biosurf-api/db"
	"github.com/dgrijalva/jwt-go"
	fiber "github.com/gofiber/fiber/v2"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func HandleLogin(c *fiber.Ctx) error {
	var users User
	var data LoginData
	err := c.BodyParser(&data)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	db.DB.Where("email = ?", data.Email).First(&users)
	if users.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Email not found",
		})
	}

	hashed := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(data.Password))
	if hashed != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"message": "Incorrect password",
			},
		)
	}

	clams := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(users.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := clams.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "smsuances_session",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: "none",
		Secure: true,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}