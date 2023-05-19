package auth

import (
	gin "github.com/gin-gonic/gin"
	bcrypt "golang.org/x/crypto/bcrypt"
	"net/http"
)

func HandleLogin(c *gin.Context) {
	var data LoginData
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	user, err := GetUserByEmail(data.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "User not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Incorrect password"})
		return
	}

	c.JSON(http.StatusOK, user)
}
