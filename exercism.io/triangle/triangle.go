package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene, or is not triangle shape at all.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	switch {
	case !isTriangle(a,b,c):
		k = NaT
	case a == b && a == c && b == c:
		k = Equ
	case (a == b && a != c) || (a == c && a != b) || (b == c && a != c):
		k = Iso
	default:
		k = Sca
	}

	return k
}

func isTriangle(a,b,c float64) bool{
	if math.IsNaN(a + b + c) || math.IsInf(a +b + c,1) || math.IsInf(a+b+c,0) {
		return false
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	return a <= b+c && b <= c+a && c <= a+b
}