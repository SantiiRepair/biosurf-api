package user

import (
	"net/http"
	"strings"
	"time"

	gin "github.com/gin-gonic/gin"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func HandleRegister(c *gin.Context) {
	var data RegisterData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	loc, err := time.LoadLocation("Europe/Madrid")
	date := time.Now().In(loc)

	user := &User{
		Name:      data.Name,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  string(passwordHash),
		CreatedAt: date,
		UpdatedAt: date,
	}

	if err := CreateUser(user); err != nil {
		if strings.Contains(err.Error(), "Email already in use") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
