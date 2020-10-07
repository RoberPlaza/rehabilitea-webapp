package auth

import (
	"log"
	"os"
	"time"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func getSecretKey() string {
	key := os.Getenv("SECRET_KEY")
	if len(key) == 0 {
		return "darksecret"
	}
	return key
}

// DefaultJWT constructs a default jwt middleware
func DefaultJWT() *jwt.GinJWTMiddleware {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:             []byte(getSecretKey()),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour * 2,
		IdentityKey:     "email",
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHeader,
		Authenticator:   Authenticate,
		Authorizator:    Authorize,
		Unauthorized:    Unauthorize,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
	})

	if err != nil {
		panic(err)
	}

	return middleware
}

// PayloadFunc is the function that creates the custom claims
func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*User); ok {
		return jwt.MapClaims{"email": v.Mail}
	}

	return jwt.MapClaims{}
}

// IdentityHeader identifies the user
func IdentityHeader(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &User{Mail: claims["email"].(string)}
}

// Authenticate authenticates a user
func Authenticate(c *gin.Context) (interface{}, error) {
	var user User
	var data LoginCredentials
	if err := c.ShouldBind(&data); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	if err := common.GetDatabase().First(&user, "mail = ?", data.Mail).Error; err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		otherPass, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
		log.Printf("user-pass: %s\ntarget-pass: %s\n", user.Password, otherPass)
		return nil, jwt.ErrFailedAuthentication
	}

	return &user, nil
}

// Authorize ...
func Authorize(data interface{}, c *gin.Context) bool {
	v, ok := data.(*User)
	return ok && v.Mail != ""
}

// Unauthorize ...
func Unauthorize(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  message,
	})
}
