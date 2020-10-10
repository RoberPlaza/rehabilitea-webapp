package main

import (
	"log"
	"time"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/auth"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/logging"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/progression"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := common.GetDatabase()
	database.InitEnv()

	base := gin.Default()
	userAuthGroup := base.Group("")
	profileAuthGroup := base.Group("")

	userAuth := auth.DefaultJWT()
	profileAuth := progression.DefaultJWT()

	base.GET("/refresh-token", userAuth.RefreshHandler)
	base.GET("/profile/login", profileAuth.LoginHandler)
	base.POST("/user/login", userAuth.LoginHandler)

	userAuthGroup.Use(userAuth.MiddlewareFunc())
	profileAuthGroup.Use(profileAuth.MiddlewareFunc())
	base.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		MaxAge:           time.Hour * 12,
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	auth.RegisterUserGroup(base.Group("/auth"))
	logging.RegisterEventRoutes(userAuthGroup.Group("/events"))
	logging.RegisterScoreRoutes(userAuthGroup.Group("/scores"))
	progression.RegisterGameGroup(userAuthGroup.Group("/games"))
	progression.RegisterProfileGroup(userAuthGroup.Group("/profiles"))
	progression.RegisterCreationGroup(userAuthGroup.Group("/new"))
	progression.RegisterDifficultyGroup(userAuthGroup.Group("/difficulties"))
	progression.RegisterGetProgressionGroup(userAuthGroup.Group("/progression"))
	progression.RegisterPostProgressionGroup(profileAuthGroup.Group("/progression"))

	log.Fatal(base.Run())
}
