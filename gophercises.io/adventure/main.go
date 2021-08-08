// Package main is main entry point for Adventure App - Choose Your Own Adventure Application
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/veselinovstf/exercises/GoLangExercises/gophercises.io/adventure/models"
)

func main() {
	// TODO
	// - reformat
	//  - Update http handling
	// - Update templating
	// - add navigation on user selection
	// - reformat

	// Parse Json to struct
	jsonAdventure, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	var arcBoard models.ArcBoard
	err = json.Unmarshal(jsonAdventure, &arcBoard)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", arcBoard))

}
