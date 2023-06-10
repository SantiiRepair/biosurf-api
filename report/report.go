package report

import (
	fiber "github.com/gofiber/fiber/v2"
)

func Report(r *fiber.App) {
	r.Post("/report", HandleReport)
}
