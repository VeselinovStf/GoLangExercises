// Package protein implements proteins functionalities
package protein

import "errors"

var (
	// ErrStop indicates terminating codons UAA, UAG, UGA
	ErrStop = errors.New("")

	// ErrInvalidBase indicates invalid base
	ErrInvalidBase = errors.New("")
)

// FromCodon translates RNA sequences into proteins.
func FromCodon(input string) (string, error) {

	switch input {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA translates RNA sequences into proteins.
func FromRNA(input string) ([]string, error) {
	out := make([]string, 0)
	for i := 0; i < len(input); i += 3 {
		fc, err := FromCodon(string(input[i : i+3]))

		if err == ErrStop {
			break
		}

		if err != nil {
			return out, err
		}

		out = append(out, fc)
	}

	return out, nil
}
