package error_handler

import (
	"errors"
	"math"
)

func MySqrt(v float64) (float64, error) {
	if v < 0 {
		return -1, errors.New("You suck give me a native number?")
	}
	return math.Sqrt(v), nil
}
