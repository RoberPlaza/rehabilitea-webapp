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
	database.InitPostConn(&common.DatabaseConnection{
		Host:      "localhost",
		User:      "postgres",
		Schema:    "postgres",
		Password:  "postgres",
		EnableSSL: false,
		Port:      5432,
	})

	base := gin.Default()
	jwtAuth := auth.DefaultJWT()
	r := base.Group("")

	base.GET("/refresh-token", jwtAuth.RefreshHandler)
	base.POST("/login", jwtAuth.LoginHandler)

	r.Use(jwtAuth.MiddlewareFunc())
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
	logging.RegisterEventRoutes(r.Group("/events"))
	logging.RegisterScoreRoutes(r.Group("/scores"))
	progression.RegisterGameGroup(r.Group("/games"))
	progression.RegisterProfileGroup(r.Group("/profiles"))
	progression.RegisterCreationGroup(r.Group("/new"))
	progression.RegisterDifficultyGroup(r.Group("/difficulties"))
	progression.RegisterProgressionGroup(r.Group("/progression"))

	log.Fatal(base.Run())
}
