package main

import (
	"os"

	"github.com/SantiiRepair/biosurf-api/auth"
	"github.com/SantiiRepair/biosurf-api/report"
	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	auth.Auth(r)
	report.Report(r)

	r.Run(":" + os.Getenv("PORT"))
}
