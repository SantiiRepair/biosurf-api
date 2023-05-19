package auth

import (
	gin "github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/register", HandleRegister)
	r.POST("/login", HandleLogin)
}
