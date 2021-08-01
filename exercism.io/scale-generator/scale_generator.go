//
package scale

var notes = []string{"C", "D", "E", "F", "G", "A", "B"}

var fiftFourth = map[string]string{
	"C" : "#",
	"F" : "b",
}

func Scale(tonic, interval string) (scale []string) {

	if interval == "" {
		return diatonicRun(tonic, fiftFourth[tonic])
	}

	return
}

func diatonicRun(note string, sign string) []string {
	var res = make([]string, 0)
	start := indexOf(note)
	for i := start; i < len(notes); i++ {
		note := stepForward(i)
		 if i >= start + 1 && sign == "b" {
			if (note != "C") && (note != "F") {
				b := note + "b"
				res = append(res, b)
			}
		}	
		res = append(res, note)

		if sign == "#" {
			if (note != "B") && (note != "E") {
				note += "#"
				res = append(res, note)
			}
		}
	}

	return res
}

func indexOf(s string) int {
	for i, v := range notes {
		if v == s {
			return i
		}
	}

	return -1
}

func stepForward(index int) string {
	if index > len(notes) {
		index = index % len(notes)
	}

	return notes[index]
}
