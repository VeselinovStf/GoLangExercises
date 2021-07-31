package proverb

import "fmt"

const (
	endProverb = "And all for the want of a %s."
	proverb    = "For want of a %s the %s was lost."
)

// Proverb generates relevant proverb in format: "For want of a <rhyme[i]> the <rhyme[i+1]> was lost."
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	result := make([]string, 0)

	for i := 0; i < len(rhyme)-1; i++ {
		result = append(result, fmt.Sprintf(proverb, rhyme[i], rhyme[i+1]))
	}

	result = append(result, fmt.Sprintf(endProverb, rhyme[0]))

	return result
}
