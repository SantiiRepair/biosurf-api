package user

import (
	"time"

	db "github.com/SantiiRepair/biosurf-api/db"
	fiber "github.com/gofiber/fiber"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func HandleRegister(c *fiber.Ctx) {
	var data RegisterData
	err := c.BodyParser(&data)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{
			"message": "Could not register",
		})
		return
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	loc, err := time.LoadLocation("Europe/Madrid")
	date := time.Now().In(loc)

	users := &User{
		Name:      data.Name,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  string(passwordHash),
		CreatedAt: date,
		UpdatedAt: date,
	}

	db.DB.Create(&users)

	c.JSON(data)

	return
}
