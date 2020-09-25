package main

import (
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/database"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/route"
	"github.com/gin-gonic/gin"
)

func main() {

	connection := database.Connection{
		Host:      "localhost",
		User:      "postgres",
		Schema:    "postgres",
		Password:  "postgres",
		EnableSSL: false,
		Port:      5432,
	}

	database.InitPostConn(&connection)

	r := gin.Default()

	route.AddUserRoutes("/user", r)
	route.AddGameRoutes("/game", r)
	route.AddCreationRoutes("/new", r)
	route.AddProgressionRoutes("/difficulty", r)
	route.AddEventRoutes("/event", r)
	route.AddScoreRoutes("/score", r)

	r.Run()
}
