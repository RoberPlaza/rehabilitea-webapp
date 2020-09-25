package controller

import (
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/database"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/model"
)

// GetAllEvents retrieves all the events from the database
func GetAllEvents() (events []model.Event, err error) {
	return events, database.Handler.Find(&events).Error
}

// GetUserEvents retrieves all the events from a user
func GetUserEvents(userID uint64) (events []model.Event, err error) {
	return events, database.Handler.Where("user_id = ?", userID).Find(&events).Error
}

// GetUserEventsAtGame ...
func GetUserEventsAtGame(userID uint64, gameName string) (events []model.Event, err error) {
	var game model.Game
	if err := database.Handler.Where("name = ?", gameName).First(&game).Error; err != nil {
		return events, err
	}
	return events, database.Handler.Find(&events).
		Where("user = ?", userID).
		Where("game = ?", game.ID).Error
}

// RegisterEvent logs an event to the database
func RegisterEvent(userID uint64, gameName string, eventType string) error {
	game, err := GetGameByName(gameName)
	if err != nil {
		return err
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	return SaveEvent(&model.Event{
		UserID:    uint64(user.ID),
		GameID:    uint64(game.ID),
		EventType: eventType,
	})
}

// SaveEvent saves an event to the database
func SaveEvent(event *model.Event) error {
	return database.Handler.Create(event).Error
}

// GetAllScores retrieves all the scores from the database
func GetAllScores() (scores []model.Score, err error) {
	return scores, database.Handler.Find(&scores).Error
}

// GetUserScores return thescores of a given user
func GetUserScores(userID uint64) (scores []model.Score, err error) {
	return scores, database.Handler.Where("user_id = ?", userID).Find(&scores).Error
}

// GetUserScoresAtGame ...
func GetUserScoresAtGame(userID uint64, gameName string) (scores []model.Score, err error) {
	var game model.Game
	if err := database.Handler.First(&game).Error; err != nil {
		return scores, err
	}

	return scores, database.Handler.
		Where("user_id = ?", userID).
		Where("game_id = ?", game.ID).
		Find(&scores).Error
}

// RegisterScore logs an score to the database
func RegisterScore(fails, maxFails, userID uint64, gameName string) error {
	game, err := GetGameByName(gameName)
	if err != nil {
		return err
	}

	return SaveScore(&model.Score{
		Fails:      fails,
		MaxAllowed: maxFails,
		UserID:     userID,
		GameID:     uint64(game.ID),
	})
}

// SaveScore stores the score to the database
func SaveScore(score *model.Score) error {
	return database.Handler.Create(score).Error
}
