// Package models implements ADVENTURE models
package models

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// ArcBoard of game arc's
type ArcBoard map[string]Arc

// Arc represent Adventure arc
type Arc struct {
	// Arc Title
	Title string `json:"title,omitempty"`
	// Arc story
	Story Story `json:"story,omitempty"`
	//	   Options be empty if it is the end of that
	//     particular story arc. Otherwise it will have one or
	//     more objects that represent an "option" that the
	//     reader has at the end of a story arc.
	Options []Option `json:"options,omitempty"`
}

// Story series of paragraphs, each represented as a string in a slice.",
type Story []string

// Option Represent arc-story options
type Option struct {
	// The text to render for this option.
	Text string `json:"text,omitempty"`
	// The name of the story arc to navigate to
	Arc string `json:"arc,omitempty"`
}

// Display shows content
func (a ArcBoard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	templatesPath := "templates/"
	switch r.URL.Path {
	case "/":
		templ, err := template.ParseFiles(templatesPath + "index.html")
		if err != nil {
			log.Fatal(err)
		}
		templ.Execute(rw, a["intro"])

	default:
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "%s", "Page not found!")
	}
}
