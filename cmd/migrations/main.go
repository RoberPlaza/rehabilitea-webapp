package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/auth"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/logging"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/progression"
)

var database = common.GetDatabase()
var userFile = flag.String("users", "data/users.json", " File with the data of the initial users")
var gameFile = flag.String("games", "data/games.json", "File with games data")
var difficultyFile = flag.String("difficulties", "data/difficulties.json", "File with difficulties data")
var modelsToMigrate = []interface{}{
	progression.Profile{},
	progression.Game{},
	progression.Difficulty{},
	progression.Progression{},
	logging.Event{},
	logging.Score{},
	auth.User{},
}

func initHandler() {
	if err := database.InitEnv(); err != nil {
		log.Fatal(err)
	}
	for _, model := range modelsToMigrate {
		database.AutoMigrate(model)
	}
}

func loadJSON(filePath string) (result []map[string]interface{}, err error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return result, err
	}

	json.Unmarshal([]byte(jsonBytes), &result)

	return result, err
}

func insertJSON() {
	if games, err := loadJSON(*gameFile); err == nil {
		for _, game := range games {
			gameName := string(game["name"].(string))
			if err := common.GetDatabase().Create(&progression.Game{Name: gameName}).Error; err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Game %s has been inserted\n", gameName)
			}
		}
	} else {
		log.Fatal(err)
	}

	if difficulties, err := loadJSON(*difficultyFile); err == nil {
		for _, difficulty := range difficulties {
			if err := common.GetDatabase().Create(&progression.Difficulty{Name: string(difficulty["name"].(string))}).Error; err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Difficulty %s inserted\n", difficulty["name"])
			}
		}
	} else {
		log.Fatal(err)
	}

	if users, err := loadJSON(*userFile); err == nil {
		for _, user := range users {
			password, _ := bcrypt.GenerateFromPassword([]byte(user["password"].(string)), bcrypt.MinCost)
			if err := common.GetDatabase().Create(
				&auth.User{
					Username: user["username"].(string),
					Mail:     user["email"].(string),
					Password: password,
				}).Error; err != nil {
				log.Fatal(err)
			} else {
				log.Printf("User %s inserted\n", user["username"])
			}
		}
	}
}

func main() {
	flag.Parse()

	log.Printf("Starting migration")
	start := time.Now()
	initHandler()
	log.Printf("Migration completed in %s", time.Since(start))

	log.Printf("Loading Json")
	start = time.Now()
	insertJSON()
	log.Printf("Initial data loaded in %s", time.Since(start))

	common.GetDatabase().Create(&progression.Profile{})
}
