package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/controller"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/database"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/model"
)

var gameFile = flag.String("games", "data/games.json", "File with games data")
var difficultyFile = flag.String("difficulties", "data/difficulties.json", "File with difficulties data")
var handler = database.Handler
var modelsToMigrate = []interface{}{
	model.Profile{},
	model.Game{},
	model.Difficulty{},
	model.Event{},
	model.Progression{},
	model.Score{},
}

func initHandler() {
	connection := database.Connection{
		Host:      "localhost",
		User:      "postgres",
		Schema:    "postgres",
		Password:  "postgres",
		Port:      5432,
		EnableSSL: false,
	}

	if err := database.InitPostConn(&connection); err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(); err != nil {
		log.Fatal(err)
	}

}

func loadJSON() {
	var games []model.Game
	var difficulties []model.Difficulty

	if file, err := ioutil.ReadFile(*gameFile); err == nil {
		if json.Unmarshal([]byte(file), &games) != nil {
			panic(fmt.Sprintf("%v", file))
		}

		for _, game := range games {
			if controller.NewGame(&game) != nil {
				panic(game)
			}
		}
	}

	if file, err := ioutil.ReadFile(*difficultyFile); err == nil {
		if json.Unmarshal([]byte(file), &difficulties) != nil {
			panic(fmt.Sprintf("%v", file))
		}

		for _, difficulty := range difficulties {
			if controller.NewDifficulty(&difficulty) != nil {
				panic(difficulty)
			}
		}
	}
	controller.NewProfile(&model.Profile{})
}

func main() {
	flag.Parse()

	log.Printf("Starting migration")
	start := time.Now()
	initHandler()
	log.Printf("Migration complete in %s", time.Since(start))

	log.Printf("Loading Json")
	start = time.Now()
	loadJSON()
	log.Printf("Initial data loaded in %s", time.Since(start))
}
