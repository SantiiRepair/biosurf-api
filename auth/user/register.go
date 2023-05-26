package user

import (
	gin "github.com/gin-gonic/gin"
	bcrypt "golang.org/x/crypto/bcrypt"
	"net/http"
)

func HandleRegister(c *gin.Context) {
    var data RegisterData
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := User{
		Email:    data.Email,
		Password: string(passwordHash),
	}
	err = CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}