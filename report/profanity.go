package report

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func HandleReport(c *gin.Context) {
	texto := c.PostForm("texto")

	if isProfanity(texto) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The text contains an obscene word"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "The text does not contain obscene words"})
}
