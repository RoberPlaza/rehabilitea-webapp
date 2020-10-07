package main

import (
	"log"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/auth"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/logging"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/progression"
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
	auth := auth.DefaultJWT()
	r := base.Group("/")

	base.GET("/refresh-token", auth.RefreshHandler)
	base.POST("/login", auth.LoginHandler)

	r.Use(auth.MiddlewareFunc())

	logging.RegisterEventRoutes(r.Group("/events"))
	logging.RegisterScoreRoutes(r.Group("/scores"))
	progression.RegisterGameGroup(r.Group("/games"))
	progression.RegisterProfileGroup(r.Group("/profiles"))
	progression.RegisterCreationGroup(r.Group("/new"))
	progression.RegisterDifficultyGroup(r.Group("/difficulties"))
	progression.RegisterProgressionGroup(r.Group("/progression"))

	log.Fatal(base.Run())
}
