// Package tabledriven demonstrates table-driven testing in Go.
// Table-driven tests are a Go best practice for testing multiple scenarios.
package tabledriven

import (
	"errors"
	"fmt"
)

// Calculator errors
var (
	ErrDivisionByZero  = errors.New("division by zero")
	ErrInvalidOperator = errors.New("invalid operator")
)

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers.
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers.
// Returns an error if attempting to divide by zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// Calculate performs a calculation based on the operator.
// Supported operators: +, -, *, /
// Returns the result and any error that occurred.
func Calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return Add(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		return Divide(a, b)
	default:
		return 0, fmt.Errorf("%w: %s", ErrInvalidOperator, operator)
	}
}
