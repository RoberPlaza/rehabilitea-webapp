package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/common"
	"github.com/RoberPlaza/rehabilitea-webapp/pkg/progression"
)

var database = common.GetDatabase()
var gameFile = flag.String("games", "data/games.json", "File with games data")
var difficultyFile = flag.String("difficulties", "data/difficulties.json", "File with difficulties data")
var modelsToMigrate = []interface{}{
	progression.Profile{},
	progression.Game{},
	progression.Difficulty{},
	progression.Progression{},
}

func initHandler() {
	connection := common.DatabaseConnection{
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
	for _, model := range modelsToMigrate {
		database.AutoMigrate(model)
	}
}

func loadJSON() {
	var games []map[string]string
	var difficulties []map[string]string

	if file, err := ioutil.ReadFile(*gameFile); err == nil {
		if json.Unmarshal([]byte(file), &games) != nil {
			panic(fmt.Sprintf("%v", file))
		}

		for _, game := range games {
			if err := progression.NewGame(&progression.Game{Name: game["name"]}); err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Game %s inserted\n", game["name"])
			}
		}
	}

	if file, err := ioutil.ReadFile(*difficultyFile); err == nil {
		if err := json.Unmarshal([]byte(file), &difficulties); err != nil {
			log.Fatal(err)
		}

		for _, difficulty := range difficulties {
			if err := progression.NewDifficulty(&progression.Difficulty{Name: difficulty["name"]}); err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Difficulty %s inserted\n", difficulty["name"])
			}
		}
	}
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

	progression.NewProfile(&progression.Profile{})
}
