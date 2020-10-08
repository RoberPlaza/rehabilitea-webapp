package auth

import (
	"github.com/gin-gonic/gin"
)

// RegisterUserGroup ...
func RegisterUserGroup(signup *gin.RouterGroup) {
	signup.POST("/signup", CreateUser)
}
