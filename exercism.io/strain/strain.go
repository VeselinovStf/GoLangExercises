// Package strain implements functions on Ints, List, Strings Collections
package strain

// Ints type of int slice
type Ints []int

// Lists type of int slice of slices
type Lists [][]int

// Strings type of string slice
type Strings []string

// Keep returns a new collection containing those elements where the predicate is true
func (i Ints) Keep(f func(int) bool) (out Ints) {
	if !validate(i) {
		return nil
	}

	for _, e := range i {
		if f(e) {
			out = append(out, e)
		}
	}

	return
}

// Discard returns a new collection containing those elements where the predicate is false
func (i Ints) Discard(f func(int) bool) Ints {
	return i.Keep(func(elem int) bool { return !f(elem) })
}

// Keep returns a new collection containing those elements where the predicate is true
func (l Lists) Keep(f func([]int) bool) (out Lists) {
	for _, e := range l {
		if f(e) {
			out = append(out, e)
		}
	}

	return
}

// Keep returns a new collection containing those elements where the predicate is true
func (s Strings) Keep(f func(string) bool) (out Strings) {
	for _, e := range s {
		if f(e) {
			out = append(out, e)
		}
	}

	return
}

func validate(v []int) bool {
	if len(v) == 0 {
		return false
	}

	return true
}
