# Debugging Go Code

## Overview
Debugging is an essential skill for any programmer. Go provides excellent tools and techniques for finding and fixing bugs in your code. This guide covers various debugging approaches, from simple print statements to advanced debugger usage.

## General Debugging Tips

### 1. Start Simple
Before reaching for a debugger, try these basic techniques:

#### Print Debugging
The simplest form of debugging - add print statements to track execution:
```go
package main

import "fmt"

func calculateSum(a, b int) int {
    fmt.Printf("DEBUG: calculateSum called with a=%d, b=%d\n", a, b)
    result := a + b
    fmt.Printf("DEBUG: result=%d\n", result)
    return result
}

func main() {
    sum := calculateSum(5, 3)
    fmt.Printf("Sum: %d\n", sum)
}
```

#### Log Package
For production-ready logging:
```go
package main

import (
    "log"
    "os"
)

func main() {
    // Set up logging with timestamps and line numbers
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    
    log.Println("Application started")
    
    value := 42
    log.Printf("Processing value: %d\n", value)
    
    if value > 100 {
        log.Fatal("Value too large!") // Logs and exits
    }
}
```

### 2. Use Go's Built-in Tools

#### Race Detector
Detect data races in concurrent code:
```bash
go test -race ./...
go run -race main.go
```

Example:
```go
package main

import (
    "fmt"
    "sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    counter++ // Race condition!
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go increment(&wg)
    }
    wg.Wait()
    fmt.Println("Counter:", counter)
}
```

Run with: `go run -race main.go` to detect the race.

#### go vet
Static analysis to find common mistakes:
```bash
go vet ./...
```

Catches issues like:
- Printf format string mismatches
- Unreachable code
- Incorrect struct tags
- Shadowed variables

## Delve Debugger

Delve is Go's dedicated debugger with powerful features for inspecting running programs.

### Installation
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

### Basic Commands

#### Starting Delve
```bash
# Debug a program
dlv debug main.go

# Debug with arguments
dlv debug main.go -- arg1 arg2

# Attach to running process
dlv attach <pid>

# Debug a test
dlv test
```

#### Essential Delve Commands
```
break (b) <location>     # Set breakpoint
    b main.main          # Break at function
    b main.go:10         # Break at line
    
continue (c)            # Continue execution
next (n)                # Step over
step (s)                # Step into
stepout                 # Step out of function

print (p) <expr>        # Print variable/expression
    p myVar
    p myVar.field
    
locals                  # Show local variables
args                    # Show function arguments
goroutines              # List goroutines
goroutine <id>          # Switch to goroutine

list (l)                # Show source code
    l main.main         # List function
    
clear <n>               # Clear breakpoint
clearall                # Clear all breakpoints

restart (r)             # Restart program
exit (quit, q)          # Exit debugger
```

### Delve Example Session

Example program to debug:
```go
// main.go
package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    for i := 0; i < 5; i++ {
        result := fibonacci(i)
        fmt.Printf("fib(%d) = %d\n", i, result)
    }
}
```

Debug session:
```bash
$ dlv debug main.go
Type 'help' for list of commands.
(dlv) break main.fibonacci
Breakpoint 1 set at 0x... for main.fibonacci() ./main.go:5
(dlv) continue
> main.fibonacci() ./main.go:5 (hits goroutine(1):1 total:1)
     1: package main
     2:
     3: import "fmt"
     4:
=>   5: func fibonacci(n int) int {
     6:     if n <= 1 {
     7:         return n
     8:     }
     9:     return fibonacci(n-1) + fibonacci(n-2)
    10: }
(dlv) args
n = 0
(dlv) next
> main.fibonacci() ./main.go:6 (PC: 0x...)
     1: package main
     2:
     3: import "fmt"
     4:
     5: func fibonacci(n int) int {
=>   6:     if n <= 1 {
     7:         return n
     8:     }
     9:     return fibonacci(n-1) + fibonacci(n-2)
    10: }
(dlv) print n
0
(dlv) continue
```

## VSCode Debugging

VSCode provides excellent Go debugging integration through the Go extension.

### Setup

1. Install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go)
2. The extension will prompt to install debugging tools including Delve

### Launch Configuration

Create `.vscode/launch.json`:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${fileDirname}"
        },
        {
            "name": "Launch File",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${file}"
        },
        {
            "name": "Launch with Arguments",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}",
            "args": ["arg1", "arg2"]
        },
        {
            "name": "Test Current File",
            "type": "go",
            "request": "launch",
            "mode": "test",
            "program": "${file}"
        }
    ]
}
```

### Using VSCode Debugger

1. **Set Breakpoints**: Click in the gutter next to line numbers
2. **Start Debugging**: Press F5 or use Debug menu
3. **Debug Controls**:
   - Continue (F5)
   - Step Over (F10)
   - Step Into (F11)
   - Step Out (Shift+F11)
   - Restart (Ctrl+Shift+F5)
   - Stop (Shift+F5)

4. **Debug Views**:
   - **Variables**: Inspect local and global variables
   - **Watch**: Add expressions to monitor
   - **Call Stack**: See function call hierarchy
   - **Breakpoints**: Manage all breakpoints
   - **Debug Console**: Execute expressions

### Example Debugging Session

```go
// main.go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func processUser(u User) {
    fmt.Printf("Processing user: %s\n", u.Name)
    // Set breakpoint here
    if u.Age < 18 {
        fmt.Println("User is a minor")
    } else {
        fmt.Println("User is an adult")
    }
}

func main() {
    users := []User{
        {Name: "Alice", Age: 25},
        {Name: "Bob", Age: 17},
        {Name: "Charlie", Age: 30},
    }
    
    for _, user := range users {
        processUser(user)
    }
}
```

**Debugging workflow**:
1. Set breakpoint at line 11 (`fmt.Printf`)
2. Press F5 to start debugging
3. Inspect `u` variable in Variables view
4. Step through conditions with F10
5. Continue to next iteration with F5

## CLI Debugging Techniques

### Using pprof for Profiling

CPU profiling:
```go
package main

import (
    "fmt"
    "os"
    "runtime/pprof"
    "time"
)

func expensiveOperation() {
    for i := 0; i < 1000000; i++ {
        _ = fmt.Sprintf("Number: %d", i)
    }
}

func main() {
    f, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    expensiveOperation()
}
```

Analyze the profile:
```bash
go run main.go
go tool pprof cpu.prof
```

Memory profiling:
```go
package main

import (
    "os"
    "runtime/pprof"
)

func allocateMemory() {
    data := make([][]byte, 100)
    for i := range data {
        data[i] = make([]byte, 1024*1024) // 1MB each
    }
}

func main() {
    allocateMemory()
    
    f, _ := os.Create("mem.prof")
    pprof.WriteHeapProfile(f)
    f.Close()
}
```

### Environment Variables for Debugging

```bash
# Enable detailed GC logging
GODEBUG=gctrace=1 go run main.go

# Show schedule trace
GODEBUG=schedtrace=1000 go run main.go

# Enable all debug output
GODEBUG=gctrace=1,schedtrace=1000 go run main.go
```

### Stack Traces

Get stack traces on panic:
```go
package main

import (
    "fmt"
    "runtime/debug"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic:", r)
            fmt.Println("Stack trace:")
            debug.PrintStack()
        }
    }()
    
    causePanic()
}

func causePanic() {
    panic("Something went wrong!")
}
```

## Best Practices

### 1. Use Structured Logging
```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("user login",
    "user_id", 123,
    "ip", "192.168.1.1",
)
```

### 2. Add Context to Errors
```go
if err != nil {
    return fmt.Errorf("failed to process user %d: %w", userID, err)
}
```

### 3. Use defer for Cleanup
```go
func processFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close() // Ensures cleanup
    
    // Process file...
    return nil
}
```

### 4. Test Edge Cases
```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name    string
        a, b    int
        want    int
        wantErr bool
    }{
        {"normal", 10, 2, 5, false},
        {"division by zero", 10, 0, 0, true},
        {"negative", -10, 2, -5, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Divide(tt.a, tt.b)
            if (err != nil) != tt.wantErr {
                t.Errorf("wantErr %v, got %v", tt.wantErr, err)
            }
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Common Debugging Scenarios

### Debugging Goroutines
```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    go func() {
        for {
            fmt.Println("Goroutine running")
            time.Sleep(time.Second)
        }
    }()
    
    // Check number of goroutines
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    
    time.Sleep(5 * time.Second)
}
```

Use Delve to inspect goroutines:
```bash
dlv debug main.go
(dlv) break main.main
(dlv) continue
(dlv) goroutines
(dlv) goroutine 5  # Switch to specific goroutine
```

### Debugging Deadlocks
```go
package main

import "sync"

func main() {
    var mu sync.Mutex
    mu.Lock()
    mu.Lock() // Deadlock!
}
```

Run with: `go run main.go` - Go will detect and report deadlock.

## Resources

- [Delve Documentation](https://github.com/go-delve/delve)
- [VSCode Go Extension](https://code.visualstudio.com/docs/languages/go)
- [Debugging Go Code with GDB](https://go.dev/doc/gdb)
- [Go Blog: Profiling Go Programs](https://go.dev/blog/pprof)
- [Go Diagnostics](https://go.dev/doc/diagnostics)

## Next Steps

- Practice setting breakpoints in your code
- Try debugging a concurrent program with race detection
- Profile a real application with pprof
- Set up VSCode debugging for your projects
- Learn advanced Delve features like conditional breakpoints
