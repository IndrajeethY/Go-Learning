# Mocking and Interfaces in Go Testing

Mocking is essential when testing code that depends on external systems like databases, APIs, or email services. Learn how to use interfaces and mocks to test your code in isolation!

## What You'll Learn

- Why we need mocks in testing
- How to use interfaces for testability
- Creating mock implementations
- Verifying behavior with mocks
- Testing error scenarios

## Files in This Folder

- `notification.go` - A notification service that sends emails
- `notification_test.go` - Tests using a mock email sender

## The Problem: Testing with Dependencies

Imagine you have code that sends emails:

```go
func NotifyUser(email, message string) error {
    // This would actually send an email!
    return smtp.Send(email, message)
}
```

**Problems when testing:**
- 🚫 Tests would send real emails
- 🚫 Needs email server configuration
- 🚫 Slow (network calls)
- 🚫 Can't test error scenarios easily

## The Solution: Interfaces and Mocks

### Step 1: Define an Interface

```go
type EmailSender interface {
    Send(to, subject, body string) error
}
```

### Step 2: Use the Interface

```go
type NotificationService struct {
    emailSender EmailSender  // Depends on interface, not implementation
}

func (s *NotificationService) NotifyUser(email, message string) error {
    return s.emailSender.Send(email, "Notification", message)
}
```

### Step 3: Create a Mock

```go
type MockEmailSender struct {
    SentEmails []SentEmail  // Track what was "sent"
    ShouldFail bool          // Control behavior
}

func (m *MockEmailSender) Send(to, subject, body string) error {
    if m.ShouldFail {
        return errors.New("mock error")
    }
    m.SentEmails = append(m.SentEmails, SentEmail{to, subject, body})
    return nil
}
```

### Step 4: Test with the Mock

```go
func TestNotifyUser(t *testing.T) {
    // Use the mock instead of real email sender
    mock := &MockEmailSender{}
    service := NewNotificationService(mock)
    
    // Test the service
    err := service.NotifyUser("user@example.com", "Hello!")
    
    // Verify behavior
    if err != nil {
        t.Fatal("unexpected error")
    }
    if len(mock.SentEmails) != 1 {
        t.Error("expected 1 email to be sent")
    }
}
```

## Running the Tests

```bash
# Run all tests
go test

# Run with verbose output
go test -v

# Run a specific test
go test -run TestNotifyUser
```

## Benefits of This Approach

✅ **No real emails sent** - Tests are safe to run anywhere
✅ **Fast** - No network calls
✅ **Reliable** - No dependency on external services
✅ **Controllable** - Can test success and failure scenarios
✅ **Verifiable** - Can check exactly what would have been sent

## Key Concepts

### 1. Interfaces Enable Testing

Interfaces let you swap implementations:
- **Production**: Use real implementation (SMTP server)
- **Testing**: Use mock implementation (records calls)

### 2. Mock Records Behavior

The mock tracks what happens:

```go
type MockEmailSender struct {
    SentEmails []SentEmail  // What was sent?
    CallCount  int           // How many times called?
    ShouldFail bool          // Should it fail?
}
```

### 3. Verify in Tests

After running your code, check the mock:

```go
// Check if the right email was sent
if mock.SentEmails[0].To != "expected@example.com" {
    t.Error("sent to wrong address")
}

// Check how many times it was called
if len(mock.SentEmails) != 2 {
    t.Error("expected 2 emails")
}
```

## Testing Different Scenarios

### Success Case

```go
func TestSuccess(t *testing.T) {
    mock := &MockEmailSender{}
    service := NewNotificationService(mock)
    
    err := service.NotifyUser("user@example.com", "Hi!")
    
    if err != nil {
        t.Fatal("should not error")
    }
    // Verify the email was sent correctly...
}
```

### Error Case

```go
func TestError(t *testing.T) {
    mock := &MockEmailSender{
        ShouldFail: true,  // Make it fail
    }
    service := NewNotificationService(mock)
    
    err := service.NotifyUser("user@example.com", "Hi!")
    
    if err == nil {
        t.Fatal("should have returned an error")
    }
}
```

### Multiple Calls

```go
func TestMultiple(t *testing.T) {
    mock := &MockEmailSender{}
    service := NewNotificationService(mock)
    
    service.NotifyUser("alice@example.com", "Hi Alice!")
    service.NotifyUser("bob@example.com", "Hi Bob!")
    
    if len(mock.SentEmails) != 2 {
        t.Error("expected 2 emails")
    }
}
```

## Common Mock Patterns

### 1. Record and Verify

```go
type Mock struct {
    Calls []Call  // Record all calls
}

func (m *Mock) Method() {
    m.Calls = append(m.Calls, Call{...})
}

// In test: verify m.Calls
```

### 2. Configurable Behavior

```go
type Mock struct {
    ReturnValue int
    ReturnError error
}

func (m *Mock) Method() (int, error) {
    return m.ReturnValue, m.ReturnError
}

// In test: configure what to return
```

### 3. Selective Failures

```go
type Mock struct {
    FailOn string  // Fail for specific input
}

func (m *Mock) Send(to string) error {
    if to == m.FailOn {
        return errors.New("fail")
    }
    return nil
}
```

## Practice Exercise

Try adding a SMS notification feature:

1. Add an interface to `notification.go`:
```go
type SMSSender interface {
    SendSMS(phone, message string) error
}
```

2. Update `NotificationService`:
```go
type NotificationService struct {
    emailSender EmailSender
    smsSender   SMSSender
}

func (s *NotificationService) NotifyBySMS(phone, message string) error {
    return s.smsSender.SendSMS(phone, message)
}
```

3. Create a mock in `notification_test.go`:
```go
type MockSMSSender struct {
    SentSMS []SentSMS
}

func (m *MockSMSSender) SendSMS(phone, message string) error {
    m.SentSMS = append(m.SentSMS, SentSMS{phone, message})
    return nil
}
```

4. Write tests for the SMS functionality!

## Next Steps

- Review [Basic Tests](../basic/) if you need to review fundamentals
- Check out [Table-Driven Tests](../table-driven/) for testing patterns
- Read [Tips](../tips.md) for more testing best practices

## Key Takeaways

✅ Use interfaces to make code testable
✅ Create mocks that implement the interface
✅ Mocks record what happens during tests
✅ Verify mock behavior to ensure code correctness
✅ Test both success and failure scenarios
✅ Mocks let you test without external dependencies
