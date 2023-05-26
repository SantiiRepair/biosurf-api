package main

import (
	// "os"

	"github.com/SantiiRepair/biosurf-api/auth/user"
	"github.com/SantiiRepair/biosurf-api/report"
	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

func main() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r := gin.Default()
	r.Use(cors.New(config))

	user.Auth(r)
	report.Report(r)

	// // r.Run(":" + os.Getenv("PORT"))
	r.Run(":7070")
}
