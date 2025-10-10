package basic

import "testing"

// TestGreet tests the Greet function with a specific name.
// This demonstrates the basic structure of a Go test.
func TestGreet(t *testing.T) {
	// Call the function we want to test
	result := Greet("Alice")
	
	// Define what we expect
	expected := "Hello, Alice!"
	
	// Check if the result matches our expectation
	if result != expected {
		// If not, report an error with details
		t.Errorf("Greet(\"Alice\") = %q; expected %q", result, expected)
	}
}

// TestGreet_EmptyName tests what happens when we pass an empty string.
// It's important to test edge cases like this!
func TestGreet_EmptyName(t *testing.T) {
	result := Greet("")
	expected := "Hello, World!"
	
	if result != expected {
		t.Errorf("Greet(\"\") = %q; expected %q", result, expected)
	}
}

// TestAdd demonstrates testing a simple arithmetic function.
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	
	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// TestAdd_Negative shows testing with negative numbers.
func TestAdd_Negative(t *testing.T) {
	result := Add(-5, 3)
	expected := -2
	
	if result != expected {
		t.Errorf("Add(-5, 3) = %d; expected %d", result, expected)
	}
}

// TestIsEven demonstrates testing a boolean function.
func TestIsEven(t *testing.T) {
	// Test an even number
	if !IsEven(4) {
		t.Error("IsEven(4) should be true")
	}
	
	// Test an odd number
	if IsEven(5) {
		t.Error("IsEven(5) should be false")
	}
	
	// Test zero (zero is even)
	if !IsEven(0) {
		t.Error("IsEven(0) should be true")
	}
}
