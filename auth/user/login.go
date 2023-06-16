package user

import (
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

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = data.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"token": t,
	})
}
