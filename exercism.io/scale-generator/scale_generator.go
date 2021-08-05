// Package scale implements music scales generation
package scale

import (
	"strings"
)

var (
	// Chromatic scale with sharps
	sharpNotes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	// Chromatic scale with flats
	flatNotes = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	// Sharps circle
	circleOfFifth = []string{"C", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
	// Flats circle
	circleOfFourth = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
	// Chromatic steps
	chromatic = "mmmmmmmmmmmm"
)

// Scale generate the musical scale by given tonic, or starting note, and a set of intervals,
// starting with the tonic and following the specified interval pattern
func Scale(tonic, interval string) (scale []string) {

	// Check for chromatic and set interval
	if interval == "" {
		interval = chromatic
	}

	// #/b Slices are uppercased - Format the input tonic string
	ft := strings.Title(tonic)
	switch {
	case contains(circleOfFifth, tonic):
		// Scale with sharp
		scale = runScale(sharpNotes, interval, ft)
	case contains(circleOfFourth, tonic):
		// Scale with flats
		scale = runScale(flatNotes, interval, ft)
	default:
		// C durr / A moll - no ligature
		scale = runScale(sharpNotes, interval, ft)
	}

	return
}

func runScale(notes []string, interval string, tonic string) (r []string) {
	start := findNote(notes, tonic)
	// Append Tonic - I
	r = append(r, notes[start])

	mem := start
	for i := start; i < len(interval)+start-1; i++ {
		// Get current step
		step := interval[i-start]

		jump := 0
		switch {
		// Major sec
		case step == 'M':
			jump = 2
		// Aug
		case step == 'A':
			jump = 3
		// Minor sec
		default:
			jump = 1
		}

		bound := mem + jump
		if bound >= len(notes) {
			// Spit to start point of notes slice
			next := bound % len(notes)
			r = append(r, notes[next])
			mem = next
		} else {
			r = append(r, notes[bound])
			mem = bound
		}
	}

	return
}

func findNote(notes []string, s string) (i int) {
	for i, v := range notes {
		if v == s {
			return i
		}
	}

	return 0
}

func contains(s []string, e string) (r bool) {
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return
}
