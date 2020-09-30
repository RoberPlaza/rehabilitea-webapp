package controller

import (
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/database"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/model"
)

// GetAllProfiles returns all profile models
func GetAllProfiles() (profiles []model.Profile, err error) {
	return profiles, database.Handler.Find(&profiles).Error
}

// GetProfileByID returs the profile with a given id if possible
func GetProfileByID(id uint64) (profile model.Profile, err error) {
	return profile, database.Handler.First(&profile, id).Error
}

// NewProfile inserts the profile into the database
func NewProfile(profile *model.Profile) (err error) {
	err = database.Handler.Create(&profile).Error

	if err == nil {
		var diff model.Difficulty
		var games []model.Game

		if database.Handler.Find(&diff, "name = ?", "easy").Error != nil {
			panic(diff)
		}

		if database.Handler.Find(&games).Error != nil {
			panic(games)
		}

		for _, game := range games {
			UpdateProfileProgression(profile, &game, &diff)
		}
	}

	return err
}

// GetAllGames return all game models
func GetAllGames() (games []model.Game, err error) {
	return games, database.Handler.Find(&games).Error
}

// GetGameByName queries a name based on the name
func GetGameByName(gameName string) (game model.Game, err error) {
	return game, database.Handler.First(&game, "name = ?", gameName).Error
}

// NewGame inserts a new game into the database
func NewGame(game *model.Game) (err error) {
	return database.Handler.Create(game).Error
}

// GetAllDifficulties return all the difficulties from the database
func GetAllDifficulties() (difficulties []model.Difficulty, err error) {
	return difficulties, database.Handler.Find(&difficulties).Error
}

// NewDifficulty inserts a new difficulty into the database
func NewDifficulty(diff *model.Difficulty) (err error) {
	return database.Handler.Create(diff).Error
}

// GetProfileGameProgression returns the progression of a profile in a game
func GetProfileGameProgression(profileID uint64, gameName string) (prog model.Progression, err error) {
	var profile model.Profile
	var game model.Game

	if profile, err = GetProfileByID(profileID); err != nil {
		return prog, err
	}

	if game, err = GetGameByName(gameName); err != nil {
		return prog, err
	}

	return prog, database.Handler.
		Order("created_at desc").
		Find(&prog,
			"profile_id = ?", profile.ID,
			"game_id = ?", game.ID,
		).Error
}

// SetProfileProgression sets the difficulty of a game for a profile
func SetProfileProgression(profileID, difficultyID uint64, gameName string) (err error) {
	var profile model.Profile
	var game model.Game
	var difficulty model.Difficulty

	if profile, err = GetProfileByID(profileID); err != nil {
		return err
	}

	if game, err = GetGameByName(gameName); err != nil {
		return err
	}

	if err = database.Handler.First(&difficulty, difficultyID).Error; err != nil {
		return err
	}

	return UpdateProfileProgression(&profile, &game, &difficulty)
}

// UpdateProfileProgression sets the difficulty of a game for a profile
func UpdateProfileProgression(profile *model.Profile, game *model.Game, difficulty *model.Difficulty) error {
	return database.Handler.Create(&model.Progression{
		Profile:    profile,
		Game:       game,
		Difficulty: difficulty,
	}).Error
}
