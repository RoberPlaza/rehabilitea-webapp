package api

import (
	"net/http"
	"strconv"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/controller"
	"github.com/gin-gonic/gin"
)

// RegisterEvent ...
func RegisterEvent(c *gin.Context) {
	var userID uint64
	var gameName string
	var eventType string

	var jsonData map[string]interface{}

	if err := c.BindJSON(&jsonData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uUserID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userID = uint64(uUserID)
	gameName = c.Param("game_name")
	eventType = jsonData["event"].(string)

	if err := controller.RegisterEvent(userID, gameName, eventType); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

// GetAllEvents ...
func GetAllEvents(c *gin.Context) {
	events, err := controller.GetAllEvents()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, events)
}

// GetUserEvents ...
func GetUserEvents(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	events, err := controller.GetUserEvents(userID)
	if err == nil {
		c.JSON(http.StatusOK, events)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetUserEventsAtGame ...
func GetUserEventsAtGame(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	events, err := controller.GetUserEventsAtGame(userID, c.Param("game_name"))
	if err == nil {
		c.JSON(http.StatusOK, events)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetAllScores ...
func GetAllScores(c *gin.Context) {
	scores, err := controller.GetAllScores()
	if err == nil {
		c.JSON(http.StatusOK, scores)
	} else {
		c.Status(http.StatusForbidden)
	}
}

// GetUserScores ..
func GetUserScores(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	scores, err := controller.GetUserScores(userID)
	if err == nil {
		c.JSON(http.StatusOK, scores)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetUserScoresAtGame ...
func GetUserScoresAtGame(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	events, err := controller.GetUserScoresAtGame(userID, c.Param("game_name"))
	if err == nil {
		c.JSON(http.StatusOK, events)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// RegisterScore ...
func RegisterScore(c *gin.Context) {
	var fails uint64
	var userID uint64
	var maxAllowed uint64
	var gameName string = c.Param("game_name")

	var jsonData map[string]float64

	if err := c.BindJSON(&jsonData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if fFails, ok := jsonData["max_fails"]; ok {
		fails = uint64(fFails)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if fMaxAllowed, ok := jsonData["fails"]; ok {
		maxAllowed = uint64(fMaxAllowed)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if controller.RegisterScore(fails, maxAllowed, userID, gameName) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
