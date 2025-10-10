# Basic Testing in Go

Welcome to Go testing! This folder contains simple examples to help you understand how testing works in Go.

## What You'll Learn

- How to write your first Go test
- Basic test structure and naming conventions
- How to check if your code works correctly
- Testing different scenarios (normal cases and edge cases)

## Files in This Folder

- `hello.go` - Contains simple functions we want to test
- `hello_test.go` - Contains tests for those functions

## Understanding the Code

### hello.go

This file has three simple functions:

1. **Greet(name)** - Returns a greeting message
2. **Add(a, b)** - Adds two numbers together
3. **IsEven(n)** - Checks if a number is even

### hello_test.go

This file contains tests for each function. Every test:
- Starts with the word `Test`
- Takes `t *testing.T` as a parameter
- Calls the function being tested
- Checks if the result is correct
- Reports an error if something is wrong

## Running the Tests

Open your terminal in this folder and run:

```bash
# Run all tests in this folder
go test

# Run tests with more details
go test -v
```

You should see output like:
```
PASS
ok      example.com/go-tests/basic    0.002s
```

## Key Concepts

### 1. Test File Naming

Test files must end with `_test.go`:
- ✅ `hello_test.go` - Correct
- ❌ `hello-test.go` - Wrong
- ❌ `test_hello.go` - Wrong

### 2. Test Function Naming

Test functions must start with `Test`:
- ✅ `func TestGreet(t *testing.T)` - Correct
- ❌ `func GreetTest(t *testing.T)` - Won't run
- ❌ `func testGreet(t *testing.T)` - Won't run (lowercase)

### 3. Reporting Errors

Use these methods to report test failures:
- `t.Error()` or `t.Errorf()` - Report error but continue test
- `t.Fatal()` or `t.Fatalf()` - Report error and stop immediately

Example:
```go
if result != expected {
    t.Errorf("Expected %d, got %d", expected, result)
}
```

## Practice Exercise

Try adding your own function and test:

1. Add a new function to `hello.go`:
```go
func Multiply(a, b int) int {
    return a * b
}
```

2. Add a test to `hello_test.go`:
```go
func TestMultiply(t *testing.T) {
    result := Multiply(3, 4)
    expected := 12
    
    if result != expected {
        t.Errorf("Multiply(3, 4) = %d; expected %d", result, expected)
    }
}
```

3. Run the tests again to see if your test passes!

## Next Steps

Once you're comfortable with basic testing, check out:
- [Table-Driven Tests](../table-driven/) - Test multiple scenarios efficiently
- [Mocking](../mocking/) - Test code that depends on other components
- [Tips](../tips.md) - Additional testing tips and best practices

## Common Questions

**Q: Why do I need tests?**
A: Tests help catch bugs early and make sure your code works as expected. They also make it safer to change your code later.

**Q: How many tests should I write?**
A: At minimum, test the normal case and edge cases (like empty inputs, zero, negative numbers, etc.).

**Q: What if my test fails?**
A: Great! That means your test caught a problem. Fix your code and run the test again.
