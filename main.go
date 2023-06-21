package main

import (
	"log"

	"github.com/SantiiRepair/biosurf-api/auth/prs"
	"github.com/SantiiRepair/biosurf-api/auth/user"
	"github.com/SantiiRepair/biosurf-api/db"
	"github.com/SantiiRepair/biosurf-api/report"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Biosurf Server",
		AppName:       "Biosurf Server v1.0.0",
	})

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	user.Auth(app)
	prs.PasswordRecovery(app)
	report.Report(app)

	log.Fatal(app.Listen(":8080"))
}
