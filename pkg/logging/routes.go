package logging

import "github.com/gin-gonic/gin"

// RegisterEventRoutes ...
func RegisterEventRoutes(events *gin.RouterGroup) {
	events.GET("/", GetAllEvents)
	events.GET("/:profile_id", GetProfileEvents)
	events.GET("/:profile_id/:game_name", GetProfileEventsAtGame)

	events.POST("/:profile_id/:game_name", RegisterNewEvent)
}

// RegisterScoreRoutes ...
func RegisterScoreRoutes(score *gin.RouterGroup) {
	score.GET("/", GetAllScores)
	score.GET("/:profile_id", GetProfileScores)
	score.GET("/:profile_id/:game_name", GetProfileScoresAtGame)

	score.POST("/:profile_id/:game_name", RegisterNewScore)
}
