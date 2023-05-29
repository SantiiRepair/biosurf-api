package main

import (
	// "os"
	"github.com/SantiiRepair/biosurf-api/auth/user"
	"github.com/SantiiRepair/biosurf-api/db"
	"github.com/SantiiRepair/biosurf-api/report"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	gin "github.com/gin-gonic/gin"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	store := cookie.NewStore([]byte("secret"))

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}

	r := gin.Default()
	r.Use(cors.New(config), sessions.Sessions("session", store))

	user.Auth(r)
	report.Report(r)

	// // r.Run(":" + os.Getenv("PORT"))
	r.Run(":7070")
}
