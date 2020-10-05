// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns their product
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide a by b - returns a non-nil error on divide-by-zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

// Sqrt calculates square root of provided value, returns a non-nil error on case of negative input
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("no support for imaginary numbers")
	}
	return math.Sqrt(a), nil
}
