// Package main is main entry point for urlshort project
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/veselinovstf/exercises/GoLangExercises/gophercises.io/urlshort/handler"
)

var (
	defaultYAML = "default.yaml"
	yaml        []byte
	yamlPath    *string
)

func init() {
	yamlPath = flag.String("yaml", defaultYAML, "Read .yaml/.yml file from file system. Format: -path: url: ")

	flag.Parse()
}

func main() {
	yamlFileContent, err := readFile(yamlPath)
	if err != nil {
		panic(err)
	}
	yaml = yamlFileContent

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
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
	c, err := os.ReadFile(*yamlPath)
	if err != nil {
		return nil, err
	}
	return c, nil
}
