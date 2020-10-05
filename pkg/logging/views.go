package logging

import (
	"net/http"
	"strconv"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	"github.com/gin-gonic/gin"
)

// GetAllEvents returns information of all the events in json format
func GetAllEvents(c *gin.Context) {
	var events []Event

	if common.GetDatabase().Find(&events).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, events)
}

// GetProfileEvents returns the events of a user in json format
func GetProfileEvents(c *gin.Context) {
	var events []Event

	if common.GetDatabase().Find(&events).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, events)
}

// GetProfileEventsAtGame returs the events of a user in a game in json
func GetProfileEventsAtGame(c *gin.Context) {
	var events []Event

	if common.GetDatabase().Find(&events, "profile_id = ?", c.Param("profile_id")).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, events)
}

// RegisterNewEvent stores a new event
func RegisterNewEvent(c *gin.Context) {
	var err error
	var event Event

	if err = c.BindJSON(&event); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if event.ProfileID, err = strconv.ParseUint(c.Param("profile_id"), 10, 32); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if common.GetDatabase().Create(&event).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllScores returns all the scores in json format
func GetAllScores(c *gin.Context) {
	var scores []Score

	if common.GetDatabase().Find(&scores).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, scores)
}

// GetProfileScores returns the scores of a game profile
func GetProfileScores(c *gin.Context) {
	var scores []Score

	if common.GetDatabase().Find(&scores, "profile_id = ?", c.Param("profile_id")).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, scores)
}

// GetProfileScoresAtGame returns the scores of a profile in a game
func GetProfileScoresAtGame(c *gin.Context) {
	var scores []Score

	if common.GetDatabase().Find(&scores,
		"profile_id = ?", c.Param("profile_id"),
		"game_name = ?", c.Param("game_name")).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, scores)
}

// RegisterNewScore saves a new score in the database
func RegisterNewScore(c *gin.Context) {
	var err error
	var score Score

	score.GameName = c.Param("game_name")

	if err := c.BindJSON(&score); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if score.ProfileID, err = strconv.ParseUint(c.Param("profile_id"), 10, 32); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if common.GetDatabase().Create(&score).Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusCreated)
}
