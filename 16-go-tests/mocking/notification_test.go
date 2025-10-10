package mocking

import (
	"errors"
	"testing"
)

// MockEmailSender is a mock implementation of the EmailSender interface.
// We use this in tests instead of a real email sender.
type MockEmailSender struct {
	// SentEmails keeps track of all emails that were "sent"
	SentEmails []SentEmail
	// ShouldFail controls whether Send should return an error
	ShouldFail bool
	// FailOnEmail makes Send fail only for a specific email address
	FailOnEmail string
}

// SentEmail represents an email that was sent (or attempted).
type SentEmail struct {
	To      string
	Subject string
	Body    string
}

// Send is our mock implementation. It doesn't actually send emails,
// but records what would have been sent for verification in tests.
func (m *MockEmailSender) Send(to, subject, body string) error {
	// If we want this mock to fail, return an error
	if m.ShouldFail || (m.FailOnEmail != "" && to == m.FailOnEmail) {
		return errors.New("mock email send failed")
	}

	// Record the email that was "sent"
	m.SentEmails = append(m.SentEmails, SentEmail{
		To:      to,
		Subject: subject,
		Body:    body,
	})

	return nil
}

// TestNotificationService_NotifyUser tests sending a notification to a single user.
func TestNotificationService_NotifyUser(t *testing.T) {
	// Create a mock email sender
	mock := &MockEmailSender{}

	// Create the service with the mock
	service := NewNotificationService(mock)

	// Test the service
	err := service.NotifyUser("user@example.com", "Hello, User!")

	// Verify no error occurred
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify an email was sent
	if len(mock.SentEmails) != 1 {
		t.Fatalf("Expected 1 email to be sent, got %d", len(mock.SentEmails))
	}

	// Verify the email content
	email := mock.SentEmails[0]
	if email.To != "user@example.com" {
		t.Errorf("Expected email to 'user@example.com', got '%s'", email.To)
	}
	if email.Subject != "Notification" {
		t.Errorf("Expected subject 'Notification', got '%s'", email.Subject)
	}
	if email.Body != "Hello, User!" {
		t.Errorf("Expected body 'Hello, User!', got '%s'", email.Body)
	}
}

// TestNotificationService_NotifyUser_Error tests error handling.
func TestNotificationService_NotifyUser_Error(t *testing.T) {
	// Create a mock that will fail
	mock := &MockEmailSender{
		ShouldFail: true,
	}

	service := NewNotificationService(mock)

	// This should return an error
	err := service.NotifyUser("user@example.com", "Hello!")

	// Verify an error was returned
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	// Verify no emails were recorded (because sending failed)
	if len(mock.SentEmails) != 0 {
		t.Errorf("Expected 0 emails to be sent, got %d", len(mock.SentEmails))
	}
}

// TestNotificationService_NotifyMultipleUsers tests sending to multiple users.
func TestNotificationService_NotifyMultipleUsers(t *testing.T) {
	mock := &MockEmailSender{}
	service := NewNotificationService(mock)

	// List of users to notify
	emails := []string{
		"alice@example.com",
		"bob@example.com",
		"charlie@example.com",
	}

	err := service.NotifyMultipleUsers(emails, "Hello, everyone!")

	// Verify no error occurred
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify all emails were sent
	if len(mock.SentEmails) != 3 {
		t.Fatalf("Expected 3 emails to be sent, got %d", len(mock.SentEmails))
	}

	// Verify each email was sent to the right person
	expectedEmails := map[string]bool{
		"alice@example.com":   false,
		"bob@example.com":     false,
		"charlie@example.com": false,
	}

	for _, email := range mock.SentEmails {
		if _, exists := expectedEmails[email.To]; exists {
			expectedEmails[email.To] = true
		}
	}

	// Check that all expected emails were sent
	for email, sent := range expectedEmails {
		if !sent {
			t.Errorf("Expected email to be sent to %s", email)
		}
	}
}

// TestNotificationService_NotifyMultipleUsers_PartialFailure tests
// what happens when sending to one user fails.
func TestNotificationService_NotifyMultipleUsers_PartialFailure(t *testing.T) {
	// Create a mock that fails for a specific email
	mock := &MockEmailSender{
		FailOnEmail: "bob@example.com",
	}
	service := NewNotificationService(mock)

	emails := []string{
		"alice@example.com",
		"bob@example.com",      // This one will fail
		"charlie@example.com",
	}

	err := service.NotifyMultipleUsers(emails, "Hello!")

	// Verify an error was returned
	if err == nil {
		t.Error("Expected an error when one send fails")
	}

	// Verify only the first email was sent (before the failure)
	if len(mock.SentEmails) != 1 {
		t.Errorf("Expected 1 email to be sent before failure, got %d", len(mock.SentEmails))
	}

	// Verify it was sent to the first user
	if len(mock.SentEmails) > 0 && mock.SentEmails[0].To != "alice@example.com" {
		t.Errorf("Expected first email to be sent to alice@example.com, got %s", 
			mock.SentEmails[0].To)
	}
}

// TestNotificationService_DifferentMessages tests using table-driven approach
// with mocks.
func TestNotificationService_DifferentMessages(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		message string
	}{
		{"greeting", "user1@example.com", "Hello!"},
		{"reminder", "user2@example.com", "Don't forget!"},
		{"alert", "user3@example.com", "Important update"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Each subtest gets its own mock
			mock := &MockEmailSender{}
			service := NewNotificationService(mock)

			err := service.NotifyUser(tt.email, tt.message)

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if len(mock.SentEmails) != 1 {
				t.Fatalf("Expected 1 email, got %d", len(mock.SentEmails))
			}

			email := mock.SentEmails[0]
			if email.To != tt.email {
				t.Errorf("Expected email to %s, got %s", tt.email, email.To)
			}
			if email.Body != tt.message {
				t.Errorf("Expected message %q, got %q", tt.message, email.Body)
			}
		})
	}
}
