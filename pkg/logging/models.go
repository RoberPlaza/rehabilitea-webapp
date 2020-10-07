package logging

import (
	"time"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/progression"
)

// Event stores the information of an ingame event
type Event struct {
	CreatedAt time.Time            `gorm:"autoCreateTime;primaryKey"`
	EventType string               `gorm:"primaryKey" binding:"required" json:"event_type"`
	ProfileID uint64               `gorm:"primaryKey"`
	GameName  uint64               `gorm:"primaryKey"`
	Profile   *progression.Profile `gorm:"foreignKey:ProfileID"`
	Game      *progression.Game    `gorm:"foreignKey:GameName"`
}

// Score stores the score of a user after completing a minigame
type Score struct {
	CreatedAt  time.Time            `gorm:"autoCreateTime;primaryKey"`
	Fails      uint64               `gorm:"not null" binding:"required" json:"fails"`
	MaxAllowed uint64               `gorm:"not null" binding:"required" json:"max_allowed"`
	ProfileID  uint64               `gorm:"primaryKey"`
	GameName   string               `gorm:"primaryKey"`
	Profile    *progression.Profile `gorm:"foreignKey:ProfileID"`
	Game       *progression.Game    `gorm:"foreignKey:GameName"`
}

// TableName returns the db name for the table that stores the model "Event"
func (Event) TableName() string {
	return "events"
}

// TableName returns the db name for the table that stores the model "Score"
func (Score) TableName() string {
	return "scores"
}
