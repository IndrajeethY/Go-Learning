// Package basic demonstrates simple Go functions and their tests.
// This is a great starting point for learning Go testing.
package basic

import "fmt"

// Greet returns a greeting message for the given name.
// If name is empty, it returns a default greeting.
func Greet(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// Add performs simple addition of two integers.
// This is a basic function to demonstrate testing arithmetic operations.
func Add(a, b int) int {
	return a + b
}

// IsEven checks if a number is even.
// Returns true if the number is divisible by 2, false otherwise.
func IsEven(n int) bool {
	return n%2 == 0
}
