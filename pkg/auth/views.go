package auth

import (
	"net/http"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a user
func CreateUser(c *gin.Context) {
	var user User
	var data NewUserData

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user.Mail = data.Mail
	user.Password = hashedPassword
	user.Username = data.Username

	if common.GetDatabase().Create(&user).Error != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	c.Status(http.StatusCreated)
}
