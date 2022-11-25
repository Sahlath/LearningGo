package sqrt

import (
	"errors"
	"math"
)

// common errors
var (
	ErrNegSqrt    = errors.New("negatve number")
	ErrNoSolution = errors.New("No solution")
)

func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}

func Sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0.0, ErrNegSqrt
	}
	return math.Sqrt(val), nil
}
