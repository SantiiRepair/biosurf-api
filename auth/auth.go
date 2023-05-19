package auth

import (
	gin"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register", HandleRegister)
	r.POST("/login", HandleLogin)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
