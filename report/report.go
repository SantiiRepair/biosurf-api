package report

import (
	fiber "github.com/gofiber/fiber"
)

func Report(r *fiber.App) {
	r.Post("/report", HandleReport)
}
