# Compiling Go Programs

## Overview
Understanding how to compile Go programs is essential for development and deployment. Go's compilation model is one of its strengths - producing fast, statically-linked binaries that are easy to deploy. This guide covers the Go toolchain and various compilation options.

## Go Build Commands

### go run
Execute Go programs without creating a binary file. Perfect for development and testing.

```bash
# Run a single file
go run main.go

# Run with multiple files
go run main.go utils.go

# Run with arguments
go run main.go arg1 arg2

# Run from a package directory
go run .

# Run a specific package
go run github.com/user/package
```

Example:
```go
// main.go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Hello from go run!")
    fmt.Println("Arguments:", os.Args[1:])
}
```

Run: `go run main.go hello world`

### go build
Compile packages and dependencies into an executable binary.

```bash
# Build current directory
go build

# Build specific file(s)
go build main.go

# Build and name output
go build -o myapp

# Build for specific package
go build github.com/user/package

# Build all packages in subdirectories
go build ./...
```

Output location:
- Current directory for packages with `main`
- `$GOPATH/pkg` for libraries

Example:
```bash
# Simple build
go build -o myapp main.go

# Run the binary
./myapp
```

### go install
Compile and install packages and dependencies. Binaries go to `$GOBIN` or `$GOPATH/bin`.

```bash
# Install current package
go install

# Install specific package
go install github.com/user/package@latest

# Install with version
go install github.com/user/package@v1.2.3

# Install tool
go install golang.org/x/tools/cmd/goimports@latest
```

Directory structure:
```
$GOPATH/
├── bin/           # Executables (go install)
├── pkg/           # Compiled packages
└── src/           # Source code (legacy)
```

## Build Flags and Options

### Common Flags

#### -o (output)
Specify output filename:
```bash
go build -o myapp
go build -o bin/myapp
```

#### -v (verbose)
Print package names as they're compiled:
```bash
go build -v
```

#### -a (force rebuild)
Force rebuild of all packages:
```bash
go build -a
```

#### -n (print commands)
Print commands but don't execute:
```bash
go build -n
```

#### -x (print commands and execute)
Show all commands as they're executed:
```bash
go build -x
```

#### -race (race detector)
Enable data race detection:
```bash
go build -race
```

#### -work (preserve work directory)
Keep the temporary work directory:
```bash
go build -work
```

### Optimization Flags

#### -gcflags
Pass flags to the Go compiler:

```bash
# Disable optimizations (easier debugging)
go build -gcflags="-N -l"

# Show optimization decisions
go build -gcflags="-m"

# Multiple flags
go build -gcflags="-N -l -m"
```

Common gcflags:
- `-N`: Disable optimizations
- `-l`: Disable inlining
- `-m`: Print optimization decisions
- `-S`: Print assembly listing

Example - See what gets inlined:
```bash
go build -gcflags="-m" main.go
```

#### -ldflags
Pass flags to the linker:

```bash
# Strip debug info (smaller binary)
go build -ldflags="-s -w"

# Set version information
go build -ldflags="-X main.version=1.2.3"

# Multiple flags
go build -ldflags="-s -w -X main.version=1.2.3"
```

Common ldflags:
- `-s`: Omit symbol table
- `-w`: Omit DWARF debug info
- `-X`: Set string variable value
- `-linkmode`: Set link mode (internal/external)

Example - Inject version at compile time:
```go
// main.go
package main

import "fmt"

var (
    version = "dev"
    commit  = "unknown"
    date    = "unknown"
)

func main() {
    fmt.Printf("Version: %s\n", version)
    fmt.Printf("Commit: %s\n", commit)
    fmt.Printf("Built: %s\n", date)
}
```

Build with version info:
```bash
go build -ldflags="-X main.version=1.0.0 -X main.commit=$(git rev-parse HEAD) -X main.date=$(date -u +%Y-%m-%d)"
```

### Build Modes

#### -buildmode
Specify the type of object file to build:

```bash
# Default executable
go build -buildmode=exe

# C shared library
go build -buildmode=c-shared -o lib.so

# C archive
go build -buildmode=c-archive -o lib.a

# Plugin
go build -buildmode=plugin -o plugin.so
```

Common build modes:
- `exe`: Executable (default)
- `c-shared`: C shared library (.so, .dll)
- `c-archive`: C archive (.a)
- `shared`: Shared library
- `plugin`: Go plugin

## Cross-Compilation

Go makes cross-compilation trivial using `GOOS` and `GOARCH` environment variables.

### Target Platforms

```bash
# List all supported platforms
go tool dist list

# Common targets:
# linux/amd64, linux/arm64, linux/386
# darwin/amd64, darwin/arm64
# windows/amd64, windows/386
# freebsd/amd64
```

### Cross-Compile Examples

#### Linux
```bash
# Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o app-linux-amd64

# Linux ARM64 (Raspberry Pi 4)
GOOS=linux GOARCH=arm64 go build -o app-linux-arm64

# Linux ARM (Raspberry Pi older)
GOOS=linux GOARCH=arm GOARM=7 go build -o app-linux-arm
```

#### macOS
```bash
# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o app-darwin-amd64

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o app-darwin-arm64
```

#### Windows
```bash
# Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o app-windows.exe

# Windows 32-bit
GOOS=windows GOARCH=386 go build -o app-windows-32.exe
```

### Build Script for Multiple Platforms

```bash
#!/bin/bash
# build-all.sh

package="myapp"
platforms=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    
    output_name=$package'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo "Building $output_name..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$output_name
    
    if [ $? -ne 0 ]; then
        echo "Build failed for $platform"
        exit 1
    fi
done

echo "All builds completed successfully!"
```

## Build Tags

Build tags control conditional compilation.

### Using Build Tags

```go
//go:build linux
// +build linux

package main

import "fmt"

func main() {
    fmt.Println("This only builds on Linux")
}
```

```go
//go:build windows
// +build windows

package main

import "fmt"

func main() {
    fmt.Println("This only builds on Windows")
}
```

### Build Tag Syntax

```go
// Build on Linux OR Darwin
//go:build linux || darwin

// Build on Linux AND amd64
//go:build linux && amd64

// Build everywhere EXCEPT Windows
//go:build !windows

// Complex conditions
//go:build (linux || darwin) && amd64
```

### Custom Build Tags

```go
//go:build integration
// +build integration

package main_test

import "testing"

func TestIntegration(t *testing.T) {
    // Integration test that requires external services
}
```

Build with custom tag:
```bash
go test -tags=integration
```

Common custom tags:
- `integration`: Integration tests
- `e2e`: End-to-end tests
- `debug`: Debug builds
- `prod`: Production builds

### File Naming Conventions

Instead of build tags, you can use file names:

```
myapp_linux.go      # Only built on Linux
myapp_windows.go    # Only built on Windows
myapp_darwin.go     # Only built on macOS
myapp_amd64.go      # Only built on amd64
myapp_arm64.go      # Only built on ARM64
```

Example:
```go
// config_linux.go
package main

const defaultPath = "/etc/myapp/config"
```

```go
// config_windows.go
package main

const defaultPath = "C:\\ProgramData\\myapp\\config"
```

## Build Constraints

### CGO

Enable or disable CGO:
```bash
# Disable CGO (static binary)
CGO_ENABLED=0 go build

# Enable CGO
CGO_ENABLED=1 go build
```

Static binary with no dependencies:
```bash
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w -extldflags "-static"' -o app
```

### Module-Aware Mode

```bash
# Enable modules
GO111MODULE=on go build

# Disable modules
GO111MODULE=off go build

# Auto (default)
GO111MODULE=auto go build
```

## Build Optimization

### Reduce Binary Size

```bash
# Strip debug info
go build -ldflags="-s -w"

# With UPX compression (requires upx tool)
go build -ldflags="-s -w" -o app
upx --brute app
```

Before and after:
```bash
# Normal build
go build -o app
ls -lh app
# -rwxr-xr-x 1 user user 6.2M Oct 10 10:00 app

# Optimized build
go build -ldflags="-s -w" -o app
ls -lh app
# -rwxr-xr-x 1 user user 4.1M Oct 10 10:00 app
```

### Faster Compilation

```bash
# Use build cache (default)
go build

# Skip tests during build
go build -tags=skip_tests

# Parallel compilation (default based on CPU cores)
go build -p 8
```

## Build Caching

Go caches build artifacts to speed up rebuilds:

```bash
# View cache location
go env GOCACHE

# Clean build cache
go clean -cache

# Clean module cache
go clean -modcache

# Clean everything
go clean -cache -modcache -testcache
```

## Practical Examples

### Development Build
Quick iteration, debugging enabled:
```bash
go build -gcflags="-N -l" -o app-dev
```

### Production Build
Optimized, small binary:
```bash
CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=$(git describe --tags)" -o app
```

### Docker Build
Static binary for minimal Docker image:
```bash
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w -extldflags "-static"' -o app
```

Dockerfile:
```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w' -o app

# Runtime stage
FROM scratch
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
```

### Makefile Example

```makefile
.PHONY: build clean test run

BINARY_NAME=myapp
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-s -w -X main.version=${VERSION} -X main.buildTime=${BUILD_TIME}"

build:
	go build ${LDFLAGS} -o ${BINARY_NAME}

build-all:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-linux-amd64
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-darwin-amd64
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-windows.exe

clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -rf bin/

test:
	go test -v ./...

run: build
	./${BINARY_NAME}

install:
	go install ${LDFLAGS}
```

Usage:
```bash
make build          # Build for current platform
make build-all      # Build for all platforms
make test           # Run tests
make run            # Build and run
make clean          # Clean build artifacts
```

## Best Practices

### 1. Use Modules
Always use Go modules for dependency management:
```bash
go mod init github.com/user/project
go mod tidy
go mod vendor  # Optional: vendor dependencies
```

### 2. Version Your Builds
Inject version info at build time:
```bash
go build -ldflags "-X main.version=$(git describe --tags)"
```

### 3. Optimize for Production
Use appropriate flags for production builds:
```bash
go build -ldflags="-s -w" -trimpath
```

### 4. Test Before Building
Always run tests before creating production builds:
```bash
go test ./... && go build
```

### 5. Use Consistent Naming
Follow naming conventions for outputs:
```bash
go build -o myapp-v1.0.0-linux-amd64
```

## Common Issues and Solutions

### Issue: Binary Too Large
**Solution**: Use ldflags to strip debug info:
```bash
go build -ldflags="-s -w"
```

### Issue: Cross-Compilation Fails with CGO
**Solution**: Disable CGO for static compilation:
```bash
CGO_ENABLED=0 GOOS=linux go build
```

### Issue: Build Takes Too Long
**Solution**: Check and clean cache if necessary:
```bash
go clean -cache
go build -p 8  # Increase parallelism
```

### Issue: Import Errors
**Solution**: Update dependencies:
```bash
go mod download
go mod tidy
```

## Resources

- [Go Command Documentation](https://pkg.go.dev/cmd/go)
- [Go Build Constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [Cross-Compilation Guide](https://go.dev/wiki/GoArm)
- [Go Compiler Directives](https://pkg.go.dev/cmd/compile)
- [Smaller Go Binaries](https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/)

## Next Steps

- Create a build script for your project
- Experiment with cross-compilation
- Try different optimization flags
- Set up automated builds with CI/CD
- Learn about build tags for different environments
