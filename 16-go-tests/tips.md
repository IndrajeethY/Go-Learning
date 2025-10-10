# Go Testing Tips and Best Practices

Additional tips, tricks, and best practices for testing in Go.

## Table of Contents
- [Test Coverage](#test-coverage)
- [Running Tests Efficiently](#running-tests-efficiently)
- [Organizing Tests](#organizing-tests)
- [Testing Best Practices](#testing-best-practices)
- [CI/CD Integration](#cicd-integration)
- [Common Patterns](#common-patterns)
- [Debugging Tests](#debugging-tests)

## Test Coverage

### Measuring Coverage

```bash
# Show coverage percentage
go test -cover

# Generate coverage profile
go test -coverprofile=coverage.out

# View coverage in HTML
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Show coverage by function
go test -coverprofile=coverage.out
go tool cover -func=coverage.out

# Coverage for all packages
go test -coverprofile=coverage.out ./...
```

### Coverage Goals

- **Aim for 80%+** for critical code
- **Don't obsess over 100%** - focus on meaningful tests
- **Test edge cases** more than happy paths
- **Some code doesn't need tests** (e.g., simple getters/setters)

### Example Output

```
$ go test -cover
PASS
coverage: 87.5% of statements
ok      mypackage    0.012s
```

## Running Tests Efficiently

### Basic Commands

```bash
# Run all tests in current package
go test

# Run all tests in all packages
go test ./...

# Verbose output (shows each test)
go test -v

# Run specific test
go test -run TestUserCreate

# Run tests matching pattern
go test -run TestUser

# Run specific subtest
go test -run TestUser/valid_email
```

### Speed Up Tests

```bash
# Run tests in parallel (faster)
go test -parallel 4

# Run only short tests (skip slow ones)
go test -short

# Cache test results (default)
go test

# Disable cache (force re-run)
go test -count=1

# Stop on first failure
go test -failfast
```

### Test Timing

```bash
# Show test duration
go test -v

# Set timeout for all tests
go test -timeout 30s

# Profile test performance
go test -cpuprofile cpu.prof
go test -memprofile mem.prof
```

## Organizing Tests

### File Organization

```
mypackage/
├── user.go           # Implementation
├── user_test.go      # Unit tests (same package)
├── export_test.go    # Test helpers for internal testing
└── integration_test.go  # Integration tests
```

### Package Names

```go
// Same package - can test private functions
package user

func TestInternalFunction(t *testing.T) { }

// Black box testing - can only test exported functions  
package user_test

func TestPublicAPI(t *testing.T) { }
```

### Test Helpers

```go
// Use t.Helper() to mark helper functions
func assertNil(t *testing.T, got interface{}) {
    t.Helper()  // Errors will point to the caller
    if got != nil {
        t.Errorf("expected nil, got %v", got)
    }
}
```

## Testing Best Practices

### 1. Use t.Run for Subtests

```go
func TestUser(t *testing.T) {
    t.Run("Create", func(t *testing.T) {
        // Test creation
    })
    
    t.Run("Update", func(t *testing.T) {
        // Test update
    })
}
```

### 2. Test Error Cases

```go
func TestDivide(t *testing.T) {
    // Test success
    result, err := Divide(10, 2)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    
    // Test error case
    _, err = Divide(10, 0)
    if err == nil {
        t.Error("expected error for division by zero")
    }
}
```

### 3. Use Table-Driven Tests

```go
func TestValidate(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantErr bool
    }{
        {"valid", "test@example.com", false},
        {"invalid", "notanemail", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := Validate(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("wantErr %v, got %v", tt.wantErr, err)
            }
        })
    }
}
```

### 4. Clean Up Resources

```go
func TestDatabase(t *testing.T) {
    db := setupDB(t)
    
    // Option 1: defer
    defer db.Close()
    
    // Option 2: t.Cleanup (runs even if test fails early)
    t.Cleanup(func() {
        db.Close()
    })
}
```

### 5. Skip Tests When Needed

```go
func TestIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test in short mode")
    }
    
    if os.Getenv("CI") == "" {
        t.Skip("skipping test in non-CI environment")
    }
}
```

### 6. Parallel Tests

```go
func TestParallel(t *testing.T) {
    tests := []struct {
        name string
        data int
    }{
        {"test1", 1},
        {"test2", 2},
    }
    
    for _, tt := range tests {
        tt := tt  // Capture range variable!
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()  // Run in parallel
            // Test code
        })
    }
}
```

## CI/CD Integration

### GitHub Actions

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out
```

### GitLab CI

```yaml
test:
  image: golang:1.21
  script:
    - go test -v -race -coverprofile=coverage.out ./...
    - go tool cover -func=coverage.out
```

### Testing Checklist for CI

```bash
# Run these before pushing
go fmt ./...           # Format code
go vet ./...           # Check for common mistakes
go test ./...          # Run all tests
go test -race ./...    # Check for race conditions
go test -cover ./...   # Check coverage
```

## Common Patterns

### Testing HTTP Handlers

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()
    
    handler(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("got %d, want %d", w.Code, http.StatusOK)
    }
}
```

### Testing with Context

```go
func TestWithContext(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := DoSomething(ctx)
    if err != nil {
        t.Fatal(err)
    }
}
```

### Testing Time-Dependent Code

```go
// Use a time provider interface
type TimeProvider interface {
    Now() time.Time
}

// Mock for testing
type MockTime struct {
    CurrentTime time.Time
}

func (m *MockTime) Now() time.Time {
    return m.CurrentTime
}
```

### Golden Files

```go
func TestOutput(t *testing.T) {
    got := generateOutput()
    
    goldenFile := "testdata/output.golden"
    
    if *update {
        // Update golden file
        os.WriteFile(goldenFile, []byte(got), 0644)
    }
    
    want, _ := os.ReadFile(goldenFile)
    
    if got != string(want) {
        t.Errorf("output mismatch")
    }
}
```

## Debugging Tests

### Print Debug Info

```go
func TestDebug(t *testing.T) {
    result := compute()
    t.Logf("result: %v", result)  // Only shown if test fails or -v flag
    
    if result != expected {
        t.Errorf("got %v, want %v", result, expected)
    }
}
```

### Use Debugger

```bash
# Use Delve debugger
dlv test

# Set breakpoint and run
(dlv) break mypackage.TestFunc
(dlv) continue
```

### Verbose Output

```bash
# Show all test output
go test -v

# Show test names as they run
go test -v ./...
```

## Advanced Testing

### Benchmarks

```go
func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Function()
    }
}

// Run with: go test -bench=.
```

### Example Tests

```go
func ExampleFunction() {
    result := Function(5)
    fmt.Println(result)
    // Output: 25
}
```

### Fuzzing (Go 1.18+)

```go
func FuzzParse(f *testing.F) {
    f.Add("test input")
    f.Fuzz(func(t *testing.T, input string) {
        Parse(input)  // Should not panic
    })
}

// Run with: go test -fuzz=.
```

## Quick Reference

### Essential Commands

```bash
go test              # Run tests
go test -v           # Verbose
go test -run Name    # Run specific test
go test -short       # Skip slow tests
go test -race        # Race detector
go test -cover       # Coverage
go test -bench=.     # Benchmarks
```

### Test Function Types

```go
func TestXxx(t *testing.T)      # Test
func BenchmarkXxx(b *testing.B) # Benchmark
func ExampleXxx()               # Example
func FuzzXxx(f *testing.F)      # Fuzz test (Go 1.18+)
```

### Assertion Methods

```go
t.Error()   / t.Errorf()    # Report error, continue
t.Fatal()   / t.Fatalf()    # Report error, stop test
t.Log()     / t.Logf()      # Log message
t.Skip()    / t.Skipf()     # Skip test
t.Helper()                  # Mark function as helper
t.Cleanup(func)             # Register cleanup
t.Parallel()                # Run test in parallel
```

## Resources

- [Go Testing Package](https://pkg.go.dev/testing)
- [Effective Go - Testing](https://go.dev/doc/effective_go#testing)
- [Table Driven Tests](https://go.dev/wiki/TableDrivenTests)
- [Go Test Comments](https://github.com/golang/go/wiki/CodeReviewComments#tests)

## Summary

✅ Write tests early and often
✅ Use table-driven tests for multiple scenarios  
✅ Test both success and error cases
✅ Keep tests fast and focused
✅ Use meaningful test names
✅ Clean up resources properly
✅ Integrate tests into CI/CD
✅ Aim for good coverage, but focus on quality

Happy testing! 🚀
