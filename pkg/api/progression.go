package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/controller"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllProfiles returns the profiles as json in a gin context
func GetAllProfiles(c *gin.Context) {
	profiles, err := controller.GetAllProfiles()
	if err == nil {
		c.JSON(http.StatusOK, profiles)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetProfileByID returs the profile with id as json in a gin context
func GetProfileByID(c *gin.Context) {
	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if profile, err := controller.GetProfileByID(profileID); err == nil {
		c.JSON(http.StatusOK, profile)
	} else {
		c.Status(http.StatusNotFound)
	}
}

// NewProfile creates a new profile and returns it as Json in a gin context
func NewProfile(c *gin.Context) {
	var profile model.Profile
	if controller.NewProfile(&profile) == nil {
		c.JSON(http.StatusOK, profile)
	} else {
		c.AbortWithStatus(http.StatusConflict)
	}
}

// GetAllGames retrieves all games from the database
func GetAllGames(c *gin.Context) {
	games, err := controller.GetAllGames()
	if err == nil {
		c.JSON(http.StatusOK, games)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetGameByName retrives the game with a certain name
func GetGameByName(c *gin.Context) {
	game, err := controller.GetGameByName(c.Param("game_name"))
	if err == nil {
		c.JSON(http.StatusOK, game)
	} else {
		c.Status(http.StatusNotFound)
	}
}

// NewGame creates a new game
func NewGame(c *gin.Context) {
	var game model.Game
	if c.BindJSON(&game) == nil {
		if err := controller.NewGame(&game); err == nil {
			c.JSON(http.StatusOK, game)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetAllDifficulties gets all the difficulties as json
func GetAllDifficulties(c *gin.Context) {
	if diff, err := controller.GetAllDifficulties(); err == nil {
		log.Println(diff)
		c.JSON(http.StatusOK, diff)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// NewDifficulty creates a new difficulty
func NewDifficulty(c *gin.Context) {
	var difficulty model.Difficulty
	if c.BindJSON(&difficulty) != nil {
		if err := controller.NewDifficulty(&difficulty); err == nil {
			c.JSON(http.StatusOK, difficulty)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetProfileProgression ...
func GetProfileProgression(c *gin.Context) {
	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if prog, err := controller.GetProfileGameProgression(profileID, c.Param("game_name")); err == nil {
		c.JSON(http.StatusOK, prog)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// SetProfileProgression ...
func SetProfileProgression(c *gin.Context) {
	var profileID, difficultyID uint64
	var jsonData map[string]interface{}
	var gameName string = c.Param("game_name")

	if err := c.BindJSON(&jsonData); err != nil {
		log.Println("Failed json parsing")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	profileID, err := strconv.ParseUint(c.Param("profile_id"), 10, 32)
	difficultyID = uint64(jsonData["difficulty"].(float64))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := controller.SetProfileProgression(profileID, difficultyID, gameName); err == nil {
		c.Status(http.StatusAccepted)
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusForbidden)
		}
	}
}
