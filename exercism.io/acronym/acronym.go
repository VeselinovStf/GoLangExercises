// Package acronym deals with acronyms manipulations
package acronym

import (
	"strings"
)

var separators = " -_"

// Abbreviate convert a phrase to its acronym.
func Abbreviate(s string) (acronym string) {
	strsplit := splitString(s, separators)

	for _, w := range strsplit {
		acronym += strings.ToUpper(string(w[0]))
	}

	return
}

func splitString(s string, sep string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(sep, r)
	})
}
