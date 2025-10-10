package tabledriven

import (
	"errors"
	"testing"
)

// TestAdd demonstrates a simple table-driven test.
// Instead of writing separate tests, we define test cases in a table.
func TestAdd(t *testing.T) {
	// Define our test cases in a slice of structs
	tests := []struct {
		name     string // Description of this test case
		a        int    // First input
		b        int    // Second input
		expected int    // Expected result
	}{
		// Each test case is easy to read and add
		{name: "positive numbers", a: 2, b: 3, expected: 5},
		{name: "negative numbers", a: -2, b: -3, expected: -5},
		{name: "mixed signs", a: 5, b: -3, expected: 2},
		{name: "with zero", a: 0, b: 5, expected: 5},
		{name: "both zero", a: 0, b: 0, expected: 0},
	}

	// Run through each test case
	for _, tt := range tests {
		// t.Run creates a subtest for each case
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSubtract shows another table-driven test example.
func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{name: "positive result", a: 5, b: 3, expected: 2},
		{name: "negative result", a: 3, b: 5, expected: -2},
		{name: "subtract zero", a: 5, b: 0, expected: 5},
		{name: "from zero", a: 0, b: 5, expected: -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestDivide demonstrates testing error cases with table-driven tests.
func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		expected  int
		wantErr   bool      // Should this test case produce an error?
		errType   error     // What type of error should it be?
	}{
		{name: "normal division", a: 10, b: 2, expected: 5, wantErr: false, errType: nil},
		{name: "division resulting in zero", a: 0, b: 5, expected: 0, wantErr: false, errType: nil},
		{name: "division by zero", a: 10, b: 0, expected: 0, wantErr: true, errType: ErrDivisionByZero},
		{name: "negative numbers", a: -10, b: 2, expected: -5, wantErr: false, errType: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			// Check if we got an error when we expected one (or didn't expect one)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide(%d, %d) error = %v; wantErr = %v", tt.a, tt.b, err, tt.wantErr)
				return
			}

			// If we expected a specific error type, check it
			if tt.wantErr && !errors.Is(err, tt.errType) {
				t.Errorf("Divide(%d, %d) error = %v; expected error type %v", tt.a, tt.b, err, tt.errType)
				return
			}

			// Only check the result if we didn't expect an error
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestCalculate demonstrates a comprehensive table-driven test with multiple operators.
// This is a realistic example of how you'd test a real function.
func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		operator string
		expected int
		wantErr  bool
	}{
		// Addition tests
		{name: "add positive", a: 2, b: 3, operator: "+", expected: 5, wantErr: false},
		{name: "add negative", a: -2, b: -3, operator: "+", expected: -5, wantErr: false},
		
		// Subtraction tests
		{name: "subtract positive", a: 5, b: 3, operator: "-", expected: 2, wantErr: false},
		{name: "subtract negative", a: 3, b: 5, operator: "-", expected: -2, wantErr: false},
		
		// Multiplication tests
		{name: "multiply positive", a: 4, b: 3, operator: "*", expected: 12, wantErr: false},
		{name: "multiply by zero", a: 5, b: 0, operator: "*", expected: 0, wantErr: false},
		{name: "multiply negative", a: -3, b: 4, operator: "*", expected: -12, wantErr: false},
		
		// Division tests
		{name: "divide evenly", a: 10, b: 2, operator: "/", expected: 5, wantErr: false},
		{name: "divide by zero", a: 10, b: 0, operator: "/", expected: 0, wantErr: true},
		
		// Error cases
		{name: "invalid operator", a: 5, b: 3, operator: "%", expected: 0, wantErr: true},
		{name: "unknown operator", a: 5, b: 3, operator: "^", expected: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.a, tt.b, tt.operator)

			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate(%d, %d, %q) error = %v; wantErr = %v", 
					tt.a, tt.b, tt.operator, err, tt.wantErr)
				return
			}

			// Check result only if no error was expected
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Calculate(%d, %d, %q) = %d; expected %d", 
					tt.a, tt.b, tt.operator, result, tt.expected)
			}
		})
	}
}

// TestMultiply is another example showing clean table structure.
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 4, 12},
		{"negative numbers", -3, -4, 12},
		{"mixed signs", -3, 4, -12},
		{"multiply by zero", 5, 0, 0},
		{"multiply by one", 7, 1, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
