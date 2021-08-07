// Package listops implements basic list opperations.
package listops

// IntList type
type IntList []int

type unaryFunc func(int) int

type predFunc func(int) bool

type binFunc func(int, int) int

// Foldl reduce each item into the accumulator from the left using function(accumulator, item))
func (l IntList) Foldl(bf func(int, int) int, a int) (out int) {
	
	for _, e := range l{
		a = bf(a,e)
	}

	return a
}

// Foldr reduce each item into the accumulator from the right using function(item, accumulator)
func (l IntList) Foldr(bf func(int, int) int, a int) (out int) {
	for _, e := range l.Reverse(){
		a = bf(e,a)
	}

	return a
}

// Reverse return a list with all the original items, but in reversed order
func (l IntList) Reverse() IntList {

	r := make([]int, len(l))
	for i := len(l) - 1; i >= 0; i-- {
		r[len(l)-1-i] = l[i]
	}

	return r
}

// Filter returns the list of all items for which predicate(item) is True
func (l IntList) Filter(pf func(int) bool) IntList {
	r := make([]int, 0)

	for _, e := range l {

		if ok := pf(e); ok {
			r = append(r, e)
		}
	}

	return r
}

// Length returns the total number of items within it
func (l IntList) Length() (count int) {
	for range l {
		count++
	}
	return
}

// Map returns the list of the results of applying function(item) on all items
func (l IntList) Map(uf func(int) int) IntList {
	r := make([]int, 0)
	for _, e := range l {
		r = append(r, uf(e))
	}

	return r
}

// Append adds all items in the second list to the end of the first list
func (l IntList) Append(in IntList) IntList {

	n := make([]int, len(l)+len(in))
	for i, e := range l {
		n[i] = e
	}

	for i, e := range in {
		n[i+len(l)] = e
	}

	return n
}

// Concat combines all items in all lists into one flattened list
func (l IntList) Concat(in []IntList) IntList {
	if len(l) == 0 {
		if len(in) == 0 {
			return l
		}

		r := make(IntList, 0)
		for _, e := range in {
			r = r.Append(e)
		}

		return r
	}

	if len(in) == 0 {
		return l
	}

	r := IntList{}
	for _, s := range in {
		t := r.Append(s)
		r = t
	}

	return l.Append(r)
}
