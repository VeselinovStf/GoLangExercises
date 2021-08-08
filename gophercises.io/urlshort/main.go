// Package main is main entry point for urlshort project
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/veselinovstf/exercises/GoLangExercises/gophercises.io/urlshort/handler"
)

var (
	defaultYAML = "default.yaml"
	defaultJSON = "default.json"
	yaml        []byte
	json        []byte
	yamlPath    *string
	jsonPath    *string
)

func init() {
	yamlPath = flag.String("yaml", defaultYAML, "Read .yaml/.yml file from file system. Format: -path: url: ")
	jsonPath = flag.String("json", defaultJSON, `Read .json file from file system. Format: [ { "path": "", "url": "" } ]`)

	flag.Parse()
}

func main() {
	yamlFileContent, err := readFile(yamlPath)
	if err != nil {
		log.Fatal(err)
	}
	yaml = yamlFileContent

	jsonFileContent, err := readFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	json = jsonFileContent

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

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readFile(filePath *string) ([]byte, error) {
	c, err := os.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}
	return c, nil
}
