package controller

import (
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/database"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/model"
)

// GetAllEvents retrieves all the events from the database
func GetAllEvents() (events []model.Event, err error) {
	return events, database.Handler.Find(&events).Error
}

// GetProfileEvents retrieves all the events from a profile
func GetProfileEvents(profileID uint64) (events []model.Event, err error) {
	return events, database.Handler.Where("profile_id = ?", profileID).Find(&events).Error
}

// GetProfileEventsAtGame ...
func GetProfileEventsAtGame(profileID uint64, gameName string) (events []model.Event, err error) {
	var game model.Game
	if err := database.Handler.Where("name = ?", gameName).First(&game).Error; err != nil {
		return events, err
	}
	return events, database.Handler.Find(&events).
		Where("profile = ?", profileID).
		Where("game = ?", game.ID).Error
}

// RegisterEvent logs an event to the database
func RegisterEvent(profileID uint64, gameName string, eventType string) error {
	game, err := GetGameByName(gameName)
	if err != nil {
		return err
	}

	profile, err := GetProfileByID(profileID)
	if err != nil {
		return err
	}

	return SaveEvent(&model.Event{
		ProfileID: uint64(profile.ID),
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

// GetProfileScores return thescores of a given profile
func GetProfileScores(profileID uint64) (scores []model.Score, err error) {
	return scores, database.Handler.Where("profile_id = ?", profileID).Find(&scores).Error
}

// GetProfileScoresAtGame ...
func GetProfileScoresAtGame(profileID uint64, gameName string) (scores []model.Score, err error) {
	var game model.Game
	if err := database.Handler.First(&game).Error; err != nil {
		return scores, err
	}

	return scores, database.Handler.
		Where("profile_id = ?", profileID).
		Where("game_id = ?", game.ID).
		Find(&scores).Error
}

// RegisterScore logs an score to the database
func RegisterScore(fails, maxFails, profileID uint64, gameName string) error {
	game, err := GetGameByName(gameName)
	if err != nil {
		return err
	}

	return SaveScore(&model.Score{
		Fails:      fails,
		MaxAllowed: maxFails,
		ProfileID:  profileID,
		GameID:     uint64(game.ID),
	})
}

// SaveScore stores the score to the database
func SaveScore(score *model.Score) error {
	return database.Handler.Create(score).Error
}
