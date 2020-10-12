package progression

import (
	"os"
	"time"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		MaxRefresh:      0,
		IdentityKey:     "id",
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
	if v, ok := data.(*Profile); ok {
		return jwt.MapClaims{"id": v.ID}
	}

	return jwt.MapClaims{}
}

// IdentityHeader identifies the user
func IdentityHeader(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &Profile{Model: gorm.Model{ID: claims["id"].(uint)}}
}

// Authenticate authenticates a user
func Authenticate(c *gin.Context) (interface{}, error) {
	var profile Profile
	var lastSession Session

	if err := c.ShouldBind(&profile); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	if err := common.GetDatabase().First(&profile).Error; err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if err := common.GetDatabase().Find(&lastSession, "profile_id = ?", profile.ID, "").Order("day desc").Error; err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	lastYear, lastMonth, lastDay := lastSession.Day.Date()
	currentYear, currentMonth, currentDay := time.Now().Date()

	if lastYear == currentYear && lastMonth == currentMonth && lastDay == currentDay {
		return nil, jwt.ErrExpiredToken
	}

	return &profile, nil
}

// Authorize ...
func Authorize(data interface{}, c *gin.Context) bool {
	v, ok := data.(*Profile)
	return ok && v.ID > 0
}

// Unauthorize ...
func Unauthorize(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  message,
	})
}
