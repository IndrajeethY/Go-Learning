# Table-Driven Testing in Go

Table-driven tests are a powerful Go testing pattern that lets you test many scenarios with minimal code duplication. This is the recommended way to write tests in Go!

## What You'll Learn

- How to organize test cases in a table
- Running multiple test scenarios efficiently
- Testing both success and error cases
- Using subtests with `t.Run()`

## Files in This Folder

- `calculator.go` - A simple calculator with basic operations
- `calculator_test.go` - Table-driven tests for the calculator

## Why Table-Driven Tests?

Instead of writing this:

```go
func TestAdd_PositiveNumbers(t *testing.T) { /* test code */ }
func TestAdd_NegativeNumbers(t *testing.T) { /* test code */ }
func TestAdd_WithZero(t *testing.T) { /* test code */ }
// ... many more similar functions
```

You write this:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"with zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; expected %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

## Running the Tests

```bash
# Run all tests
go test

# Run with verbose output (shows each subtest)
go test -v

# Run a specific test
go test -run TestCalculate

# Run a specific subtest
go test -run TestCalculate/add_positive
```

Example output:
```
=== RUN   TestCalculate
=== RUN   TestCalculate/add_positive
=== RUN   TestCalculate/subtract_positive
=== RUN   TestCalculate/divide_by_zero
--- PASS: TestCalculate (0.00s)
    --- PASS: TestCalculate/add_positive (0.00s)
    --- PASS: TestCalculate/subtract_positive (0.00s)
    --- PASS: TestCalculate/divide_by_zero (0.00s)
```

## Understanding the Pattern

### 1. Define Test Cases

Create a slice of structs with your test cases:

```go
tests := []struct {
    name     string  // Description of the test case
    // ... input fields
    expected int     // Expected result
    wantErr  bool    // Should this case produce an error?
}{
    {name: "test case 1", /* ... */},
    {name: "test case 2", /* ... */},
}
```

### 2. Loop Through Cases

Iterate over the test cases and run each one:

```go
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // Run the test with tt.input
        // Check against tt.expected
    })
}
```

### 3. Benefits

- **Easy to add new cases** - Just add a line to the table
- **Less code duplication** - Testing logic written once
- **Clear test names** - Each subtest has a descriptive name
- **Easy to spot patterns** - See all test cases at a glance

## Testing Error Cases

When testing functions that return errors, add a `wantErr` field:

```go
tests := []struct {
    name     string
    input    int
    wantErr  bool  // Do we expect an error?
}{
    {"valid input", 5, false},
    {"invalid input", -1, true},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        result, err := MyFunction(tt.input)
        
        // Check if error matches expectation
        if (err != nil) != tt.wantErr {
            t.Errorf("MyFunction() error = %v, wantErr %v", err, tt.wantErr)
            return
        }
        
        // Only check result if no error was expected
        if !tt.wantErr {
            // check result...
        }
    })
}
```

## Practice Exercise

Try adding a modulo operation to the calculator:

1. Add to `calculator.go`:
```go
func Modulo(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivisionByZero
    }
    return a % b, nil
}
```

2. Add a table-driven test to `calculator_test.go`:
```go
func TestModulo(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
        wantErr  bool
    }{
        {"normal modulo", 10, 3, 1, false},
        {"modulo by zero", 10, 0, 0, true},
        {"even division", 10, 5, 0, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Modulo(tt.a, tt.b)
            if (err != nil) != tt.wantErr {
                t.Errorf("Modulo(%d, %d) error = %v; wantErr = %v", 
                    tt.a, tt.b, err, tt.wantErr)
                return
            }
            if !tt.wantErr && result != tt.expected {
                t.Errorf("Modulo(%d, %d) = %d; expected %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

## Common Patterns

### Compact Table Format

For simple cases, you can use a more compact format:

```go
tests := []struct {
    name     string
    a, b     int
    expected int
}{
    {"case 1", 1, 2, 3},
    {"case 2", 5, 3, 8},
}
```

### Verbose Table Format

For complex cases with many fields:

```go
tests := []struct {
    name     string
    a        int
    b        int
    expected int
}{
    {
        name:     "case 1",
        a:        1,
        b:        2,
        expected: 3,
    },
}
```

## Next Steps

- Check out [Mocking](../mocking/) to learn how to test code with dependencies
- Read [Tips](../tips.md) for more testing best practices
- Explore the [Basic Tests](../basic/) if you need to review fundamentals

## Key Takeaways

✅ Table-driven tests reduce code duplication
✅ Use `t.Run()` to create subtests with descriptive names
✅ Add a `wantErr` field when testing error cases
✅ The table format makes it easy to add new test cases
✅ Each subtest runs independently and reports separately
