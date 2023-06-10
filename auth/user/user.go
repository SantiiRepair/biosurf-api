package user

import (
	fiber "github.com/gofiber/fiber"
)

const SecretKey = "secret"

func Auth(r *fiber.App) {
	r.Post("/user/register", HandleRegister)
	r.Post("/user/login", HandleLogin)
	r.Post("/user/logout", HandleLogout)
}
