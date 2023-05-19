package auth

import (
	gin "github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	r.POST("/register", HandleRegister)
	r.POST("/login", HandleLogin)
}
