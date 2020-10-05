package auth

import (
	"net/http"
	"time"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	"github.com/dgrijalva/jwt-go"
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

// Login autehnticates a user
func Login(c *gin.Context) {
	var user User
	var loginData LoginCredentials

	if err := c.BindJSON(&loginData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if common.GetDatabase().First(&user, "email = ?", loginData.Mail).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(loginData.Password)); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Status(http.StatusAccepted)
}

// GenerateToken generates a token for the user
func GenerateToken(user *User) (string, error) {
	claims := UserClaims{
		Mail: user.Mail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	return token.SignedString([]byte( /*os.Getenv("SECRET_KEY") */ "darksecret"))
}

func ValidateToken(encodedToken string) error {
	return nil
}
