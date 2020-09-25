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

// GetAllUsers returns the users as json in a gin context
func GetAllUsers(c *gin.Context) {
	users, err := controller.GetAllUsers()
	if err == nil {
		c.JSON(http.StatusOK, users)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetUserByID returs the user with id as json in a gin context
func GetUserByID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if user, err := controller.GetUserByID(userID); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.Status(http.StatusNotFound)
	}
}

// NewUser creates a new user and returns it as Json in a gin context
func NewUser(c *gin.Context) {
	var user model.User
	if controller.NewUser(&user) == nil {
		c.JSON(http.StatusOK, user)
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

// GetUserProgression ...
func GetUserProgression(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if prog, err := controller.GetUserGameProgression(userID, c.Param("game_name")); err == nil {
		c.JSON(http.StatusOK, prog)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// SetUserProgression ...
func SetUserProgression(c *gin.Context) {
	var userID, difficultyID uint64
	var jsonData map[string]interface{}
	var gameName string = c.Param("game_name")

	if err := c.BindJSON(&jsonData); err != nil {
		log.Println("Failed json parsing")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	difficultyID = uint64(jsonData["difficulty"].(float64))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := controller.SetUserProgression(userID, difficultyID, gameName); err == nil {
		c.Status(http.StatusAccepted)
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusForbidden)
		}
	}
}
