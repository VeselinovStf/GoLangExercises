// Package gigasecond implements opperations on time with 10^9 (1,000,000,000) seconds intervals
package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

// AddGigasecond determine the moment that would be after a gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(time.Duration(math.Pow(10, 9)) * time.Second))
}
