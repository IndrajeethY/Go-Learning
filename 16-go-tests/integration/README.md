# Integration Testing

Integration tests verify that different components of your system work together correctly. Unlike unit tests that test individual functions in isolation, integration tests check how multiple parts interact.

## What Are Integration Tests?

Integration tests:
- Test multiple components working together
- May involve databases, APIs, or file systems
- Are typically slower than unit tests
- Provide confidence that the whole system works

## Coming Soon!

This folder will contain examples of:

### Database Integration Tests
Testing code that interacts with databases:
```go
func TestUserRepository_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }
    
    // Set up test database
    db := setupTestDB(t)
    defer db.Close()
    
    // Test database operations
    repo := NewUserRepository(db)
    // ... test code
}
```

### HTTP API Integration Tests
Testing HTTP handlers and routes:
```go
func TestAPIEndpoints_Integration(t *testing.T) {
    // Start test server
    server := httptest.NewServer(handler)
    defer server.Close()
    
    // Make requests and verify responses
    resp, err := http.Get(server.URL + "/users")
    // ... test code
}
```

### External Service Integration Tests
Testing code that calls external APIs or services:
```go
func TestExternalAPI_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }
    
    // Call real external API (or test version)
    client := NewAPIClient()
    result, err := client.FetchData()
    // ... test code
}
```

## Best Practices for Integration Tests

### 1. Use Build Tags

Separate integration tests from unit tests:
```go
//go:build integration
// +build integration

package mypackage

func TestIntegration(t *testing.T) {
    // Integration test code
}
```

Run with: `go test -tags=integration`

### 2. Use testing.Short()

Skip slow tests when running quick tests:
```go
func TestIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test in short mode")
    }
    // Test code
}
```

Run quick tests: `go test -short`

### 3. Set Up and Tear Down

Always clean up resources:
```go
func TestDatabase(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    // Or use t.Cleanup()
    t.Cleanup(func() {
        db.Close()
    })
}
```

### 4. Use Test Containers

For database tests, consider using Docker containers:
```go
func setupTestDB(t *testing.T) *sql.DB {
    // Start a Docker container with a test database
    // Run migrations
    // Return connection
}
```

### 5. Isolate Tests

Each integration test should:
- Start with a known state
- Not depend on other tests
- Clean up after itself

## Example Structure

```
myproject/
├── internal/
│   └── user/
│       ├── user.go
│       ├── user_test.go           # Unit tests
│       └── user_integration_test.go  # Integration tests
```

## Running Integration Tests

```bash
# Run only unit tests (skip integration)
go test -short ./...

# Run only integration tests
go test -run Integration ./...

# Run with build tags
go test -tags=integration ./...

# Run with longer timeout
go test -timeout 30s ./...
```

## When to Write Integration Tests

Write integration tests when:
- ✅ Testing database operations
- ✅ Testing HTTP APIs
- ✅ Testing file I/O
- ✅ Testing external service integrations
- ✅ Testing multiple components together

Keep unit tests for:
- ✅ Business logic
- ✅ Pure functions
- ✅ Fast, isolated tests

## Resources

- [Go Blog - Integration Tests](https://go.dev/doc/tutorial/add-a-test)
- [Testcontainers for Go](https://golang.testcontainers.org/)
- [httptest Package](https://pkg.go.dev/net/http/httptest)

## Contributing

Feel free to contribute integration test examples to this folder! Good examples would include:
- SQLite in-memory database tests
- HTTP handler tests with httptest
- File system operation tests
- Mock external API tests

Check back soon for complete examples!
