// Package darts implements darts game functionalities
package darts

import "math"

type circle struct {
	radius float64
	score  int
}

var darts = []circle{
	{radius: 1.0, score: 10},
	{radius: 5.0, score: 5},
	{radius: 10.0, score: 1},
}

// Score returns returns the correct amount earned by a dart landing in that point
// by  given point in the target (defined by its real cartesian coordinates x and y).
func Score(x, y float64) (score int) {

	for _, c := range darts {
		if s := circleInCartesianPlane(x, y, c.radius); s <= c.radius {
			return c.score
		}
	}

	return
}

func circleInCartesianPlane(x, y float64, r float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
