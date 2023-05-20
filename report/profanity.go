package report

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func HandleReport(c *gin.Context) {
	text := c.PostForm("text")

	if isProfanity(text) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The text contains an obscene word"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "The text does not contain obscene words"})
}
