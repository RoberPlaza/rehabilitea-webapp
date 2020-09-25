package route

import (
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/api"
	"github.com/gin-gonic/gin"
)

// AddUserRoutes ...
func AddUserRoutes(path string, group *gin.Engine) {
	users := group.Group(path)

	users.GET("/", api.GetAllUsers)
	users.GET("/:user_id", api.GetUserByID)
}

// AddGameRoutes ...
func AddGameRoutes(path string, group *gin.Engine) {
	games := group.Group(path)

	games.GET("/", api.GetAllGames)
	games.GET("/:game_name", api.GetGameByName)
}

// AddProgressionRoutes adds the progression path to a gin group
func AddProgressionRoutes(path string, group *gin.Engine) {
	progressions := group.Group(path)

	progressions.GET("/", api.GetAllDifficulties)

	progressions.GET("/:user_id/:game_name", api.GetUserProgression)
	progressions.POST("/:user_id/:game_name", api.SetUserProgression)
}

// AddCreationRoutes ...
func AddCreationRoutes(path string, group *gin.Engine) {
	creations := group.Group(path)

	creations.GET("/user", api.NewUser)
	creations.GET("/game", api.NewGame)
	creations.GET("/difficulty", api.NewDifficulty)

	creations.POST("/user", api.NewUser)
	creations.POST("/game", api.NewGame)
	creations.POST("/difficulty", api.NewDifficulty)
}

// AddEventRoutes ...
func AddEventRoutes(path string, group *gin.Engine) {
	events := group.Group(path)

	events.GET("/", api.GetAllEvents)
	events.GET("/:user_id", api.GetUserEvents)
	events.GET("/:user_id/:game_name", api.GetUserEventsAtGame)

	events.POST("/:user_id/:game_name", api.RegisterEvent)
}

// AddScoreRoutes ...
func AddScoreRoutes(path string, group *gin.Engine) {
	scores := group.Group(path)

	scores.GET("/", api.GetAllScores)
	scores.GET("/:user_id", api.GetUserScores)
	scores.GET("/:user_id/:game_name", api.GetUserScoresAtGame)

	scores.POST("/:user_id/:game_name", api.RegisterScore)
}
