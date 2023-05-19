package main

import (
	"github.com/SantiiRepair/biosurf-api/auth"
	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	auth.Routes(r)

	r.Run(":8080")
}
