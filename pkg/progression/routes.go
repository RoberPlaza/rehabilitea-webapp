package progression

import "github.com/gin-gonic/gin"

// RegisterProfileGroup returns a group that handlers the game profiles
func RegisterProfileGroup(profiles *gin.RouterGroup) {
	profiles.GET("/", AllProfilesView)
	profiles.GET("/:profile_id", ProfileByIDView)
}

// RegisterGameGroup ...
func RegisterGameGroup(games *gin.RouterGroup) {
	games.GET("/", AllGamesView)
	games.GET("/:game_name", GameByNameView)
}

// RegisterGetProgressionGroup adds the progression path to a gin group
func RegisterGetProgressionGroup(progressions *gin.RouterGroup) {
	progressions.GET("/", AllDifficultiesView)
	progressions.GET("/:profile_id/:game_name", ProfileProgressionView)
}

// RegisterCreationGroup ...
func RegisterCreationGroup(creations *gin.RouterGroup) {
	creations.GET("/profile", NewProfileView)
	creations.GET("/game", NewGameView)
	creations.GET("/difficulty", NewDifficultyView)

	creations.POST("/profile", NewProfileView)
	creations.POST("/game", NewGameView)
	creations.POST("/difficulty", NewDifficultyView)
}

// RegisterDifficultyGroup ...
func RegisterDifficultyGroup(difficulties *gin.RouterGroup) {
	difficulties.GET("/", AllDifficultiesView)
	difficulties.GET("/:difficulty_name", DifficultyByNameView)
}

// RegisterPostProgressionGroup ...
func RegisterPostProgressionGroup(progressions *gin.RouterGroup) {
	progressions.POST("/:profile_id/:game_name", SetProfileProgressionView)
}
