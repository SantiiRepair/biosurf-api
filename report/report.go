package report

import (
	gin "github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	r.POST("/report", HandleReport)
}

