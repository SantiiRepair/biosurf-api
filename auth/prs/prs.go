package prs

import (
	fiber "github.com/gofiber/fiber/v2"
)

func PasswordRecovery(r *fiber.App) {
	r.Post("/recovery", HandlePRS)
}
