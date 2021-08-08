// Package main is main entry point for urlshort project
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/veselinovstf/exercises/GoLangExercises/gophercises.io/urlshort/db"
	"github.com/veselinovstf/exercises/GoLangExercises/gophercises.io/urlshort/handler"
)

var (
	serverPort  = ":8080"
	defaultYAML = "default.yaml"
	defaultJSON = "default.json"
	yaml        []byte
	json        []byte
	yamlPath    *string
	jsonPath    *string
	useDb       *bool
)

func init() {
	yamlPath = flag.String("yaml", defaultYAML, "Read .yaml/.yml file from file system. Format: -path: url: ")
	jsonPath = flag.String("json", defaultJSON, `Read .json file from file system. Format: [ { "path": "", "url": "" } ]`)
	useDb = flag.Bool("useDb", false, `Load urls from database`)

	flag.Parse()

	yaml = loadContent(yamlPath)
	json = loadContent(jsonPath)
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		log.Fatal(err)
	}

	jsonHandler, err := handler.JSONHandler(json, yamlHandler)
	if err != nil {
		log.Fatal(err)
	}

	if *useDb {
		dbConnection := db.NewConnection(
			"DESKTOP-QK2ADMV",
			"URLShortner",
			"chofex",
			"",
		)

		dbHandler, err := handler.SQLHandler(dbConnection, jsonHandler)
		if err != nil {
			log.Fatal(err)
		}
		startServer(dbHandler)
	} else {
		startServer(jsonHandler)
	}

}

func startServer(handler http.HandlerFunc) {
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(serverPort, handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	return mux
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to URLShort!")
}

func readFile(filePath *string) ([]byte, error) {
	c, err := os.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func loadContent(filePath *string) []byte {
	content, err := readFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return content
}
