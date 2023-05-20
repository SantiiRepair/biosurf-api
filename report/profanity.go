package report

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func HandleReport(c *gin.Context) {
	text := c.PostForm("text")
	// id := c.PostForm("id")

	if isProfanity(text) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The text contains an obscene word"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The text does not contain obscene words"})
}
