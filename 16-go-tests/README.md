# Go Testing Tutorial for Beginners

Welcome! This is a beginner-friendly guide to testing in Go. If you're new to Go testing, you're in the right place! 🎉

## Why Testing Matters

Testing helps you:
- ✅ Catch bugs before they reach production
- ✅ Make changes confidently (tests catch breaking changes)
- ✅ Document how your code works
- ✅ Write better, more reliable code

Go makes testing easy with its built-in `testing` package - no external tools needed!

## Tutorial Structure

This tutorial is organized into folders, each teaching a specific testing concept. Start from the beginning and work your way through:

### 1️⃣ [Basic Testing](./basic/)
**Start here!** Learn the fundamentals of Go testing.

**What you'll learn:**
- How to write your first test
- Test file naming conventions (`_test.go`)
- Using `t.Errorf()` to report failures
- Testing different functions

**Files:**
- `hello.go` - Simple functions to test
- `hello_test.go` - Your first tests
- `README.md` - Detailed explanation

**Quick start:**
```bash
cd basic/
go test -v
```

### 2️⃣ [Table-Driven Testing](./table-driven/)
Learn the Go way of testing multiple scenarios efficiently.

**What you'll learn:**
- The table-driven testing pattern
- Testing multiple cases with one function
- Using `t.Run()` for subtests
- Testing error conditions

**Files:**
- `calculator.go` - Calculator with multiple operations
- `calculator_test.go` - Table-driven tests
- `README.md` - Detailed explanation

**Quick start:**
```bash
cd table-driven/
go test -v
```

### 3️⃣ [Mocking and Interfaces](./mocking/)
Test code that depends on external services or components.

**What you'll learn:**
- Why we need mocks
- Using interfaces for testability
- Creating mock implementations
- Testing without real dependencies

**Files:**
- `notification.go` - Email notification service
- `notification_test.go` - Tests with mocks
- `README.md` - Detailed explanation

**Quick start:**
```bash
cd mocking/
go test -v
```

### 4️⃣ [Integration Testing](./integration/)
**Coming soon!** Test how multiple components work together.

**What you'll learn:**
- Database integration tests
- HTTP API testing
- Testing with external services

**Status:** Placeholder with examples and best practices

### 📝 [Tips and Best Practices](./tips.md)
Additional testing tips, coverage, CI/CD, and more.

**What you'll learn:**
- Measuring test coverage
- Running tests efficiently
- CI/CD integration
- Common testing patterns
- Debugging tests

## Quick Start Guide

### Prerequisites

- Go installed (version 1.16 or later)
- Basic Go knowledge (variables, functions, packages)
- A terminal/command prompt

### Running Your First Test

1. **Navigate to the basic folder:**
```bash
cd go-tests/basic/
```

2. **Run the tests:**
```bash
go test
```

You should see:
```
PASS
ok      example.com/go-tests/basic    0.002s
```

3. **Run with verbose output:**
```bash
go test -v
```

### Understanding Test Output

```
=== RUN   TestGreet
--- PASS: TestGreet (0.00s)
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      example.com/go-tests/basic    0.002s
```

- `=== RUN` - Test is starting
- `--- PASS` - Test passed
- `ok` - All tests in the package passed
- Time shows how long tests took

## Learning Path

Follow this recommended path through the tutorial:

```
1. Basic Testing (1-2 hours)
   ↓
2. Table-Driven Testing (1-2 hours)
   ↓
3. Mocking and Interfaces (2-3 hours)
   ↓
4. Tips and Best Practices (reference)
   ↓
5. Integration Testing (when ready)
```

## Key Testing Concepts

### Test Files Must:
- ✅ End with `_test.go` (e.g., `math_test.go`)
- ✅ Be in the same folder as the code they test
- ✅ Import the `testing` package

### Test Functions Must:
- ✅ Start with `Test` (e.g., `TestAdd`)
- ✅ Take one parameter: `t *testing.T`
- ✅ Use `t.Error()` or `t.Errorf()` to report failures

### Example:
```go
package basic

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
    }
}
```

## Essential Testing Commands

```bash
# Run all tests in current folder
go test

# Run with detailed output
go test -v

# Run specific test
go test -run TestAdd

# Run tests in all subfolders
go test ./...

# Run tests with coverage report
go test -cover

# Run only fast tests (skip slow integration tests)
go test -short

# Check for race conditions
go test -race
```

## Go Testing Features

### Built-in Testing Package
No need to install anything! Go includes everything you need:
```go
import "testing"
```

### Fast Execution
Go tests compile and run quickly, giving you instant feedback:
```
ok      mypackage    0.008s
```

### Parallel Testing
Run tests in parallel for even faster execution:
```go
func TestParallel(t *testing.T) {
    t.Parallel()
    // Test code
}
```

### Benchmarking
Measure performance of your code:
```go
func BenchmarkMyFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MyFunction()
    }
}
```

## Best Practices

✅ **Write tests as you code** - Don't wait until the end
✅ **Test edge cases** - Empty strings, zero values, negative numbers
✅ **Use descriptive names** - `TestUserCreate_WithInvalidEmail`
✅ **Keep tests simple** - Each test should check one thing
✅ **Use table-driven tests** - For testing multiple scenarios
✅ **Test behavior, not implementation** - Test what it does, not how

## Common Questions

**Q: Do I need to test every function?**
A: Focus on testing public functions and important logic. Simple getters/setters often don't need tests.

**Q: How much test coverage should I aim for?**
A: Aim for 80%+ coverage, but focus on quality over quantity. Test important code thoroughly.

**Q: Should I write tests before or after coding?**
A: Either works! Some prefer TDD (tests first), others prefer writing tests as they code.

**Q: How do I test code that talks to a database or API?**
A: Use mocks and interfaces! Check out the [Mocking tutorial](./mocking/).

**Q: My test is failing - what do I do?**
A: Great! Your test caught a problem. Read the error message, fix your code, and run the test again.

## Getting Help

- **Read the folder READMEs** - Each folder has detailed explanations
- **Look at the code examples** - Learn by reading and running them
- **Modify the examples** - Change them and see what happens
- **Check [tips.md](./tips.md)** - For additional patterns and tricks

## Resources

### Official Documentation
- [Go Testing Package](https://pkg.go.dev/testing)
- [Go Testing Tutorial](https://go.dev/doc/tutorial/add-a-test)
- [Effective Go - Testing](https://go.dev/doc/effective_go#testing)

### Additional Learning
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [Table Driven Tests](https://go.dev/wiki/TableDrivenTests)
- [Go Blog - Testing](https://go.dev/blog/examples)

## What's Next?

After completing this tutorial:

1. **Practice!** Add tests to your own Go projects
2. **Explore advanced topics** - Fuzzing, benchmarking, profiling
3. **Set up CI/CD** - Automate testing in your development workflow
4. **Learn integration testing** - Test how components work together
5. **Contribute!** Help improve this tutorial

## Contributing

Found a mistake? Want to add an example? Contributions are welcome!

## Summary

Testing in Go is:
- ✅ **Easy** - Built into the language
- ✅ **Fast** - Quick to run
- ✅ **Powerful** - Rich features
- ✅ **Essential** - For reliable code

Start with [Basic Testing](./basic/) and happy testing! 🚀
