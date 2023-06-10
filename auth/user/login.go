package user

import (
	db "github.com/SantiiRepair/biosurf-api/db"
	"github.com/dgrijalva/jwt-go"
	fiber "github.com/gofiber/fiber/v2"
	bcrypt "golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

func HandleLogin(c *fiber.Ctx) error {
	var data LoginData
	err := c.BodyParser(&data)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	var users User
	db.DB.Where("email = ?", users.Email).First(&users)

	if users.ID == 0 {
		c.Status(fiber.StatusNotFound)
		c.JSON(fiber.Map{
			"message": "Email not found",
		})
	}

	hashed := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(data.Password))
	if hashed != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(
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
		c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	c.Status(fiber.StatusAccepted)
	c.JSON(fiber.Map{
		"message": "success",
	})

	return err
}
