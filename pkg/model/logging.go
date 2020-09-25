package model

import (
	"time"
)

// Event stores the information of an ingame event
type Event struct {
	CreatedAt time.Time `gorm:"autoCreateTime;primaryKey"`
	EventType string    `gorm:"primaryKey" binding:"required"`
	UserID    uint64    `gorm:"primaryKey" binding:"required"`
	GameID    uint64    `gorm:"primaryKey" binding:"required"`
	User      *User     `gorm:"foreignKey:UserID"`
	Game      *Game     `gorm:"foreignKey:GameID"`
}

// Score stores the score of a user after completing a minigame
type Score struct {
	CreatedAt  time.Time `gorm:"primaryKey;primaryKey" binding:"required"`
	Fails      uint64    `gorm:"not null" binding:"required"`
	MaxAllowed uint64    `binding:"required"`
	UserID     uint64    `gorm:"primaryKey" binding:"required"`
	GameID     uint64    `gorm:"primaryKey" binding:"required"`
	User       *User     `gorm:"foreignKey:UserID"`
	Game       *Game     `gorm:"foreignKey:GameID"`
}

// TableName returns the db name for the table that stores the model "Event"
func (Event) TableName() string {
	return "events"
}

// TableName returns the db name for the table that stores the model "Score"
func (Score) TableName() string {
	return "scores"
}
