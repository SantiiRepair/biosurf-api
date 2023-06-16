package user

import (
	"time"

	db "github.com/SantiiRepair/biosurf-api/db"
	fiber "github.com/gofiber/fiber/v2"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func HandleRegister(c *fiber.Ctx) error {
	var peer User
	var data RegisterData
	err := c.BodyParser(&data)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{
			"message": "Could not register",
		})
	}

	db.DB.Where("email = ?", data.Email).First(&peer)
	if peer.ID > 0 {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "Email exists",
		})
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	loc, err := time.LoadLocation("Europe/Madrid")
	date := time.Now().In(loc)

	users := &User{
		Name:      data.Name,
		Lastname:  data.Lastname,
		Email:     data.Email,
		Password:  string(passwordHash),
		CreatedAt: date.String(),
		UpdatedAt: date.String(),
	}

	db.DB.Create(&users)

	c.JSON(data)

	return err
}
