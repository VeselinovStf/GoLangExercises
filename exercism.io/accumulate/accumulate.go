// Package accumulate implements functionality for string accumulations
package accumulate

// Accumulate perform operation on each collection element,
// returns a new collection containing the result of applying func
// operation to each element of the input collection
func Accumulate(input []string, command func(string) string) (res []string) {

	for _, v := range input {
		res = append(res, command(v))
	}

	return
}
