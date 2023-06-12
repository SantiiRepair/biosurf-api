package user

import (
	fiber "github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"

func Auth(r *fiber.App) {
	r.Post("/user/register", HandleRegister)
	r.Post("/user/login", HandleLogin)
	r.Post("/user/logout", HandleLogout)
}
