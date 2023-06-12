package mails

import (
	fiber "github.com/gofiber/fiber/v2"
)

func EmailNotifier(r *fiber.App) {
	r.Post("/notifier", HandleNotifier)
}
