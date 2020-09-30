package api

import (
	"net/http"
	"strconv"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/controller"
	"github.com/gin-gonic/gin"
)

// RegisterEvent ...
func RegisterEvent(c *gin.Context) {
	var profileID uint64
	var gameName string
	var eventType string

	var jsonData map[string]interface{}

	if err := c.BindJSON(&jsonData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uProfileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	profileID = uint64(uProfileID)
	gameName = c.Param("game_name")
	eventType = jsonData["event"].(string)

	if err := controller.RegisterEvent(profileID, gameName, eventType); err != nil {
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

// GetProfileEvents ...
func GetProfileEvents(c *gin.Context) {
	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	events, err := controller.GetProfileEvents(profileID)
	if err == nil {
		c.JSON(http.StatusOK, events)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetProfileEventsAtGame ...
func GetProfileEventsAtGame(c *gin.Context) {
	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	events, err := controller.GetProfileEventsAtGame(profileID, c.Param("game_name"))
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

// GetProfileScores ..
func GetProfileScores(c *gin.Context) {
	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	scores, err := controller.GetProfileScores(profileID)
	if err == nil {
		c.JSON(http.StatusOK, scores)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetProfileScoresAtGame ...
func GetProfileScoresAtGame(c *gin.Context) {
	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	events, err := controller.GetProfileScoresAtGame(profileID, c.Param("game_name"))
	if err == nil {
		c.JSON(http.StatusOK, events)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// RegisterScore ...
func RegisterScore(c *gin.Context) {
	var fails uint64
	var profileID uint64
	var maxAllowed uint64
	var gameName string = c.Param("game_name")

	var jsonData map[string]float64

	if err := c.BindJSON(&jsonData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)
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

	if controller.RegisterScore(fails, maxAllowed, profileID, gameName) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
