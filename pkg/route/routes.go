package route

import (
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/api"
	"github.com/gin-gonic/gin"
)

// AddProfileRoutes ...
func AddProfileRoutes(path string, group *gin.Engine) {
	profiles := group.Group(path)

	profiles.GET("/", api.GetAllProfiles)
	profiles.GET("/:profile_id", api.GetProfileByID)
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

	progressions.GET("/:profile_id/:game_name", api.GetProfileProgression)
	progressions.POST("/:profile_id/:game_name", api.SetProfileProgression)
}

// AddCreationRoutes ...
func AddCreationRoutes(path string, group *gin.Engine) {
	creations := group.Group(path)

	creations.GET("/profile", api.NewProfile)
	creations.GET("/game", api.NewGame)
	creations.GET("/difficulty", api.NewDifficulty)

	creations.POST("/profile", api.NewProfile)
	creations.POST("/game", api.NewGame)
	creations.POST("/difficulty", api.NewDifficulty)
}

// AddEventRoutes ...
func AddEventRoutes(path string, group *gin.Engine) {
	events := group.Group(path)

	events.GET("/", api.GetAllEvents)
	events.GET("/:profile_id", api.GetProfileEvents)
	events.GET("/:profile_id/:game_name", api.GetProfileEventsAtGame)

	events.POST("/:profile_id/:game_name", api.RegisterEvent)
}

// AddScoreRoutes ...
func AddScoreRoutes(path string, group *gin.Engine) {
	scores := group.Group(path)

	scores.GET("/", api.GetAllScores)
	scores.GET("/:profile_id", api.GetProfileScores)
	scores.GET("/:profile_id/:game_name", api.GetProfileScoresAtGame)

	scores.POST("/:profile_id/:game_name", api.RegisterScore)
}
