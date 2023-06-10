package main

import (
	"github.com/SantiiRepair/biosurf-api/auth/user"
	"github.com/SantiiRepair/biosurf-api/db"
	"github.com/SantiiRepair/biosurf-api/report"
	fiber "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	db.Connect()
	

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    })) 

	user.Auth(app)
	report.Report(app)

	app.Listen(":8080")
}
