// Package dna implements library for DNA processing functions
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {

	h := Histogram{
		'C': 0,
		'G': 0,
		'T': 0,
		'A': 0,
	}

	for _, v := range d {
		if _, ok := h[v]; ok {
			h[v]++
		} else {
			return Histogram{}, errors.New("can't count not existing nucleotide")
		}
	}

	return h, nil
}
