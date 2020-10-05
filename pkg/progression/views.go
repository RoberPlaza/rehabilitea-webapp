package progression

import (
	"net/http"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	"github.com/gin-gonic/gin"
)

// AllProfilesView returns the profiles as json in a gin context
func AllProfilesView(c *gin.Context) {
	var profiles []Profile

	if common.GetDatabase().Find(&profiles).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, profiles)
}

// ProfileByIDView returs the profile with id as json in a gin context
func ProfileByIDView(c *gin.Context) {
	var profile Profile

	if common.GetDatabase().First(&profile, c.Param("profile_id")).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, profile)
}

// NewProfileView creates a new profile and returns it as Json in a gin context
func NewProfileView(c *gin.Context) {
	var games []Game
	var profile Profile
	var difficulty Difficulty

	if common.GetDatabase().Find(&difficulty, "name = ?", "easy").Error != nil ||
		common.GetDatabase().Find(&games).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if common.GetDatabase().Create(&profile).Error == nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	c.JSON(http.StatusOK, profile)
}

// AllGamesView retrieves all games from the database
func AllGamesView(c *gin.Context) {
	var games []Game

	if common.GetDatabase().Find(&games).Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, games)
}

// GameByNameView retrives the game with a certain name
func GameByNameView(c *gin.Context) {
	var game Game

	if common.GetDatabase().Find(&game, "name = ?", c.Param("game_name")).Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, game)
}

// NewGameView creates a new game
func NewGameView(c *gin.Context) {
	var game Game

	if c.ShouldBindJSON(&game) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if common.GetDatabase().Create(&game).Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, game)
}

// AllDifficultiesView gets all the difficulties as json
func AllDifficultiesView(c *gin.Context) {
	var difficulties []Difficulty

	if common.GetDatabase().Find(&difficulties).Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, difficulties)
}

// DifficultyByNameView ...
func DifficultyByNameView(c *gin.Context) {
	diff := Difficulty{Name: c.Param("difficulty_name")}

	if common.GetDatabase().First(&diff).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, diff)
}

// NewDifficultyView creates a new difficulty
func NewDifficultyView(c *gin.Context) {
	var difficulty Difficulty

	if c.BindJSON(&difficulty) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if common.GetDatabase().Create(&difficulty).Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, difficulty)
}

// ProfileProgressionView ...
func ProfileProgressionView(c *gin.Context) {
	var game Game
	var profile Profile
	var progression Progression

	if common.GetDatabase().First(&profile, c.Param("profile_id")).Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if common.GetDatabase().First(&game, "name = ?", c.Param("game_name")).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if common.GetDatabase().First(&progression, "profile_id = ?", profile.ID, "game_id = ?", game.ID).
		Order("created_at DESC").Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, progression)
}

// SetProfileProgressionView ...
func SetProfileProgressionView(c *gin.Context) {
	var game Game
	var profile Profile
	var difficulty Difficulty
	var jsonData map[string]interface{}

	if err := c.BindJSON(&jsonData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if common.GetDatabase().First(&profile, c.Param("profile_id")).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if common.GetDatabase().First(&game, "name = ?", c.Param("game_name")).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if common.GetDatabase().First(&difficulty, "name = ?", jsonData["difficulty"]).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if common.GetDatabase().Create(&Progression{
		Profile:    &profile,
		Game:       &game,
		Difficulty: &difficulty}).Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusAccepted)
}
