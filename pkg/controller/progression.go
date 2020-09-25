package controller

import (
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/database"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/model"
)

// GetAllUsers returns all user models
func GetAllUsers() (users []model.User, err error) {
	return users, database.Handler.Find(&users).Error
}

// GetUserByID returs the user with a given id if possible
func GetUserByID(id uint64) (user model.User, err error) {
	return user, database.Handler.First(&user, id).Error
}

// NewUser inserts the user into the database
func NewUser(user *model.User) (err error) {
	err = database.Handler.Create(&user).Error

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
			UpdateUserProgression(user, &game, &diff)
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

// GetUserGameProgression returns the progression of a user in a game
func GetUserGameProgression(userID uint64, gameName string) (prog model.Progression, err error) {
	var user model.User
	var game model.Game

	if user, err = GetUserByID(userID); err != nil {
		return prog, err
	}

	if game, err = GetGameByName(gameName); err != nil {
		return prog, err
	}

	return prog, database.Handler.
		Order("created_at desc").
		Find(&prog,
			"user_id = ?", user.ID,
			"game_id = ?", game.ID,
		).Error
}

// SetUserProgression sets the difficulty of a game for a user
func SetUserProgression(userID, difficultyID uint64, gameName string) (err error) {
	var user model.User
	var game model.Game
	var difficulty model.Difficulty

	if user, err = GetUserByID(userID); err != nil {
		return err
	}

	if game, err = GetGameByName(gameName); err != nil {
		return err
	}

	if err = database.Handler.First(&difficulty, difficultyID).Error; err != nil {
		return err
	}

	return UpdateUserProgression(&user, &game, &difficulty)
}

// UpdateUserProgression sets the difficulty of a game for a user
func UpdateUserProgression(user *model.User, game *model.Game, difficulty *model.Difficulty) error {
	return database.Handler.Create(&model.Progression{
		User:       user,
		Game:       game,
		Difficulty: difficulty,
	}).Error
}
