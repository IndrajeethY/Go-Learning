# Efficient Go Usage

## Overview
Writing efficient Go code goes beyond just making it work - it's about writing idiomatic, performant, and maintainable code that follows Go's philosophy and best practices. This guide covers essential patterns, idioms, and techniques for efficient Go programming.

## Go Idioms and Conventions

### 1. Use Short Variable Names
Go prefers short, meaningful names in limited scopes.

```go
// Good
func (s *Server) handleRequest(r *http.Request) {
    u := r.URL
    p := u.Path
    // ...
}

// Avoid
func (server *Server) handleRequest(request *http.Request) {
    url := request.URL
    path := url.Path
    // ...
}
```

**Guidelines**:
- Single letter for very short scopes: `i`, `j` for loops
- Short names for common concepts: `r` for reader, `w` for writer
- Longer names for larger scopes and exported identifiers

### 2. Accept Interfaces, Return Structs

```go
// Good - accepts interface for flexibility
func ProcessData(r io.Reader) (*Result, error) {
    // Can accept any reader: file, network, buffer, etc.
    return &Result{}, nil
}

// Avoid - overly specific
func ProcessData(f *os.File) (*Result, error) {
    return &Result{}, nil
}
```

### 3. Handle Errors Immediately

```go
// Good
func readFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("failed to open %s: %w", path, err)
    }
    defer f.Close()
    
    data, err := io.ReadAll(f)
    if err != nil {
        return nil, fmt.Errorf("failed to read %s: %w", path, err)
    }
    
    return data, nil
}

// Avoid - delayed error handling
func readFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    defer f.Close()
    
    if err != nil {  // f.Close() called on nil file!
        return nil, err
    }
    
    return io.ReadAll(f)
}
```

### 4. Use defer for Cleanup

```go
func processFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close() // Guaranteed cleanup
    
    // Process file...
    // Even if panic occurs, file will be closed
    
    return nil
}
```

### 5. Use Blank Identifier for Unused Variables

```go
// When you need to ignore a return value
_, err := doSomething()

// When implementing an interface but not using receiver
func (s *Service) method() {
    // Not using s
}

// Better: use blank identifier in parameter
func (*Service) method() {
    // Explicitly shows we don't need receiver
}
```

## Naming Best Practices

### Variables and Functions

```go
// Good names
userCount := 10
maxConnections := 100
httpClient := &http.Client{}

// Avoid
uc := 10
mc := 100
c := &http.Client{}  // Too short for package scope

// Function names should be verbs
func parseConfig() {}
func validateInput() {}
func sendEmail() {}

// Avoid
func config() {}  // Noun
func validation() {}  // Noun
```

### Package Names

```go
// Good
package user
package http
package json

// Avoid
package user_management  // No underscores
package userManagement   // No camelCase
package users            // Usually singular
```

### Exported vs Unexported

```go
package api

// Exported - starts with capital letter
type User struct {
    Name  string  // Exported field
    Email string  // Exported field
    age   int     // Unexported field
}

// Exported function
func NewUser(name, email string) *User {
    return &User{Name: name, Email: email}
}

// Unexported function
func validateEmail(email string) bool {
    return true
}
```

## Error Handling Patterns

### 1. Wrap Errors with Context

```go
package main

import (
    "fmt"
    "os"
)

func readConfig(path string) error {
    _, err := os.Open(path)
    if err != nil {
        // Wrap error with context
        return fmt.Errorf("config error: %w", err)
    }
    return nil
}

func main() {
    if err := readConfig("/etc/app.conf"); err != nil {
        fmt.Fprintf(os.Stderr, "startup failed: %v\n", err)
    }
}
```

### 2. Create Custom Error Types

```go
package main

import (
    "errors"
    "fmt"
)

// Custom error type
type ValidationError struct {
    Field string
    Value interface{}
    Err   error
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s=%v: %v", 
        e.Field, e.Value, e.Err)
}

func (e *ValidationError) Unwrap() error {
    return e.Err
}

// Usage
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field: "age",
            Value: age,
            Err:   errors.New("must be non-negative"),
        }
    }
    return nil
}
```

### 3. Use errors.Is and errors.As

```go
package main

import (
    "errors"
    "fmt"
    "os"
)

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
    return fmt.Errorf("user lookup: %w", ErrNotFound)
}

func main() {
    err := findUser(123)
    
    // Check for specific error
    if errors.Is(err, ErrNotFound) {
        fmt.Println("User not found")
    }
    
    // Extract error type
    var pathErr *os.PathError
    if errors.As(err, &pathErr) {
        fmt.Println("Path error:", pathErr.Path)
    }
}
```

### 4. Don't Panic

```go
// Bad - panics are for truly exceptional cases
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")  // Don't do this!
    }
    return a / b
}

// Good - return error
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// OK - panic for impossible conditions
func mustCompile(pattern string) *regexp.Regexp {
    re, err := regexp.Compile(pattern)
    if err != nil {
        panic(err)  // OK in init or with constant patterns
    }
    return re
}
```

## Memory Optimization

### 1. Use Pointers Wisely

```go
// Large structs - use pointers to avoid copying
type LargeStruct struct {
    data [1000]int
    // ... many fields
}

func process(s *LargeStruct) {  // Pass by pointer
    // Work with s
}

// Small structs - value is fine
type Point struct {
    X, Y int
}

func distance(p Point) float64 {  // Pass by value
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
```

### 2. Preallocate Slices When Size is Known

```go
// Good - preallocate
func processItems(count int) []Item {
    items := make([]Item, 0, count)  // Preallocate capacity
    for i := 0; i < count; i++ {
        items = append(items, Item{ID: i})
    }
    return items
}

// Less efficient - multiple allocations
func processItems(count int) []Item {
    var items []Item  // Will grow dynamically
    for i := 0; i < count; i++ {
        items = append(items, Item{ID: i})
    }
    return items
}
```

### 3. Reuse Buffers

```go
package main

import (
    "bytes"
    "sync"
)

// Pool of buffers
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func processData(data []byte) string {
    // Get buffer from pool
    buf := bufferPool.Get().(*bytes.Buffer)
    defer func() {
        buf.Reset()
        bufferPool.Put(buf)  // Return to pool
    }()
    
    buf.Write(data)
    // Process...
    return buf.String()
}
```

### 4. Use String Builder for Concatenation

```go
// Bad - creates new string each time
func buildString(parts []string) string {
    result := ""
    for _, part := range parts {
        result += part  // Allocates new string
    }
    return result
}

// Good - efficient string building
func buildString(parts []string) string {
    var sb strings.Builder
    for _, part := range parts {
        sb.WriteString(part)
    }
    return sb.String()
}
```

## Benchmarking

### Writing Benchmarks

```go
package main

import (
    "strings"
    "testing"
)

// Benchmark string concatenation
func BenchmarkStringConcat(b *testing.B) {
    parts := []string{"hello", "world", "from", "go"}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        result := ""
        for _, part := range parts {
            result += part
        }
        _ = result
    }
}

// Benchmark string builder
func BenchmarkStringBuilder(b *testing.B) {
    parts := []string{"hello", "world", "from", "go"}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        for _, part := range parts {
            sb.WriteString(part)
        }
        _ = sb.String()
    }
}
```

Run benchmarks:
```bash
go test -bench=. -benchmem

# Output shows Builder is much faster:
# BenchmarkStringConcat-8     2000000    850 ns/op   80 B/op  4 allocs/op
# BenchmarkStringBuilder-8    5000000    280 ns/op   48 B/op  1 allocs/op
```

### Benchmark Best Practices

```go
func BenchmarkExample(b *testing.B) {
    // Setup code (not measured)
    data := make([]int, 1000)
    
    // Reset timer before measured code
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        // Code to benchmark
        process(data)
    }
}

// Prevent compiler optimizations
func BenchmarkCalculation(b *testing.B) {
    var result int
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        result = expensiveCalculation()
    }
    
    // Use result to prevent optimization
    _ = result
}
```

## Profiling

### CPU Profiling

```go
package main

import (
    "os"
    "runtime/pprof"
)

func main() {
    f, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    // Your code here
    doWork()
}

func doWork() {
    for i := 0; i < 1000000; i++ {
        _ = complexCalculation(i)
    }
}
```

Analyze:
```bash
go run main.go
go tool pprof cpu.prof

# Interactive commands:
# top    - show top functions
# list funcName - show source with annotations
# web    - visualize in browser
```

### Memory Profiling

```go
package main

import (
    "os"
    "runtime/pprof"
)

func main() {
    doWork()
    
    f, _ := os.Create("mem.prof")
    pprof.WriteHeapProfile(f)
    f.Close()
}
```

### Continuous Profiling with net/http/pprof

```go
package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"  // Import for side effects
)

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Your application code
    runServer()
}
```

Access profiles:
```bash
# CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Heap profile
go tool pprof http://localhost:6060/debug/pprof/heap

# Goroutine profile
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

## Concurrency Patterns

### 1. Worker Pool

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        // Process job
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    var wg sync.WaitGroup
    
    // Start workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }
    
    // Send jobs
    for i := 1; i <= 9; i++ {
        jobs <- i
    }
    close(jobs)
    
    // Close results when all workers done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Println(result)
    }
}
```

### 2. Pipeline Pattern

```go
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    // Set up pipeline
    nums := generator(1, 2, 3, 4)
    squared := square(nums)
    
    // Consume values
    for n := range squared {
        fmt.Println(n)
    }
}
```

### 3. Fan-Out, Fan-In

```go
func fanOut(in <-chan int, n int) []<-chan int {
    channels := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        channels[i] = worker(in)
    }
    return channels
}

func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}
```

## Interface Design

### Small Interfaces

```go
// Good - small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Compose when needed
type ReadWriter interface {
    Reader
    Writer
}

// Avoid - large interfaces
type Database interface {
    Connect()
    Disconnect()
    Query()
    Insert()
    Update()
    Delete()
    Transaction()
    // ... many more methods
}
```

### Accept Minimal Interfaces

```go
// Good - accepts minimal interface
func Copy(dst Writer, src Reader) error {
    // Can work with any Reader/Writer
    return nil
}

// Avoid - overly specific
func Copy(dst *os.File, src *os.File) error {
    return nil
}
```

## Best Practices Summary

### DO:
✅ Use `go fmt` to format code  
✅ Use `go vet` to check for issues  
✅ Write tests for your code  
✅ Handle all errors explicitly  
✅ Use meaningful variable names  
✅ Keep functions small and focused  
✅ Prefer composition over inheritance  
✅ Document exported identifiers  
✅ Use interfaces for abstraction  
✅ Benchmark performance-critical code  

### DON'T:
❌ Ignore errors  
❌ Use panic for normal error handling  
❌ Create unnecessary goroutines  
❌ Prematurely optimize  
❌ Use global variables excessively  
❌ Have deeply nested code  
❌ Write long functions (>50 lines)  
❌ Use generic names like `data`, `tmp`  
❌ Expose internal details  
❌ Mix concerns in a single function  

## Code Review Checklist

- [ ] Is the code formatted with `go fmt`?
- [ ] Are all errors handled?
- [ ] Are variable names clear and concise?
- [ ] Is the code properly documented?
- [ ] Are there tests?
- [ ] Is the interface design minimal?
- [ ] Are resources properly cleaned up (defer)?
- [ ] Is concurrency used appropriately?
- [ ] Are there any race conditions?
- [ ] Is the code idiomatic Go?

## Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Go Proverbs](https://go-proverbs.github.io/)
- [Dave Cheney's Blog](https://dave.cheney.net/)

## Next Steps

- Apply these patterns to your existing code
- Write benchmarks for performance-critical sections
- Review code with these idioms in mind
- Practice profiling your applications
- Study well-written Go projects (Kubernetes, Docker, etc.)
