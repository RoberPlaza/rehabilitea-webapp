package main

import (
	"log"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
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

	r := gin.Default()

	progression.RegisterGameGroup(r.Group("/games"))
	progression.RegisterProfileGroup(r.Group("/profiles"))
	progression.RegisterCreationGroup(r.Group("/new"))
	progression.RegisterDifficultyGroup(r.Group("/difficulties"))
	progression.RegisterProgressionGroup(r.Group("/progression"))

	log.Fatal(r.Run())
}
