// Package strand implements DNA and RNA strands functional methods.
package strand

var nucleotides = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA converts given DNA strand and returns its RNA complement (per RNA transcription).
func ToRNA(dna string) (rna string) {

	for _, c := range dna {
		if val, ok := nucleotides[c]; ok {
			rna += val
		}
	}

	return
}
