package user

import (
	"github.com/gin-contrib/sessions"
	gin "github.com/gin-gonic/gin"
	bcrypt "golang.org/x/crypto/bcrypt"
	"net/http"
)

func HandleLogin(c *gin.Context) {
	var data LoginData
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := GetUserByEmail(data.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
