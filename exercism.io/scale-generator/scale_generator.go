// Package scale implements music scales generation
package scale

var sharpNotes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatNotes = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var circleOfFifth = []string{"C", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
var circleOfFourth = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

// Scale generate the musical scale by given tonic, or starting note, and a set of intervals,
// starting with the tonic and following the specified interval pattern
func Scale(tonic, interval string) (scale []string) {

	if interval == "" {
		// hromatic
		interval = "mmmmmmmmmmm"
	}

	if ok, _ := contains(circleOfFifth, tonic); ok {
		// Use sharps
		start := findNote(sharpNotes, tonic)
		scale = runScale(sharpNotes, interval, start)

		return
	}

	if ok, _ := contains(circleOfFourth, tonic); ok {
		// Use flats
		start := findNote(sharpNotes, tonic)
		scale = runScale(flatNotes, interval, start)

		return
	}

	// No armature
	start := findNote(sharpNotes, tonic)
	scale = runScale(sharpNotes, interval, start)

	return
}

func runScale(notes []string, interval string, start int) (r []string) {

	length := len(interval)

	c := 0

	for i := start; i <= length+start; i++ {

		step := interval[c]

		jump := 0
		switch {
		case step == 'M':
			jump = 1
		case step == 'A':
			jump = 2
		default:
			jump = 0
		}

		if i != start {
			if i > len(notes)+jump-1 {
				r = append(r, notes[i%length-1+jump])
			} else {
				r = append(r, notes[i+jump])
			}

			c++
		} else {
			r = append(r, notes[i])
		}

	}

	return
}

func findNote(notes []string, s string) (i int) {
	_, i = contains(notes, s)

	return
}

func contains(s []string, e string) (r bool, i int) {
	for i, v := range s {
		if v == e {
			return true, i
		}
	}

	return
}
