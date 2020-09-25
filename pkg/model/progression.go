package model

import (
	"time"

	"gorm.io/gorm"
)

// User stores the information of a game participant
type User struct {
	gorm.Model
	Birthday time.Time `json:"birthday"`
}

// Game stores the information of a minigame of the application
type Game struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null" binding:"required"`
}

// Difficulty stores the possible difficulties of a minigame
type Difficulty struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null" binding:"required"`
}

// Progression stores a step in the user progression
type Progression struct {
	CreatedAt    time.Time   `gorm:"autoCreateTime"`
	UserID       uint64      `gorm:"primaryKey"`
	GameID       uint64      `gorm:"primaryKey"`
	DifficultyID uint64      `gorm:"primaryKey"`
	User         *User       `gorm:"foreignKey:UserID"`
	Game         *Game       `gorm:"foreignKey:GameID"`
	Difficulty   *Difficulty `gorm:"foreignKey:DifficultyID"`
}

// TableName returns the db name for the table that stores the model "Game"
func (Game) TableName() string {
	return "games"
}

// TableName returns the db name for the table that stores the model "Difficulty"
func (Difficulty) TableName() string {
	return "difficulties"
}

// TableName returns the db name for the table that stores the model "Progression"
func (Progression) TableName() string {
	return "progressions"
}

// TableName returns the db name for the table that stores the model "User"
func (User) TableName() string {
	return "users"
}
