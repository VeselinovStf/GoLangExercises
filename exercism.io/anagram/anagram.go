// Package anagram implements anagram functions
package anagram

import (
	"sort"
	"strings"
)

// Detect returns sublist of anagrams of the given word from candidates
func Detect(word string, candidates []string) (out []string) {
	nword := sortString(strings.ToLower(word))
	for _, s := range candidates {
		if len(s) == len(word) && strings.ToLower(word) != strings.ToLower(s) {
			sn := sortString(strings.ToLower(s))
			anagram := false
			for i := range nword {
				if nword[i] != sn[i] {
					anagram = false
					break
				} else {
					anagram = true
				}
			}

			if anagram {
				out = append(out, s)
			}
		}
	}

	return
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
