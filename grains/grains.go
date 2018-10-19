// Package grains calculats the number of grains of wheat on a chessboard given the the number on each square doubles
package grains

import (
	"errors"
	"math"
)

// Square finds the the number of grains on chessboard
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("input should be between 1 and 64")
	}
	return uint64(math.Pow(2, float64(input-1))), nil
}

// Total returns the total grains on chessboard
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		v, _ := Square(i)
		sum += v
	}
	return sum
}
