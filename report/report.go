package report

import (
	gin "github.com/gin-gonic/gin"
)

func Report(r *gin.Engine) {
	r.POST("/report", HandleReport)
}

