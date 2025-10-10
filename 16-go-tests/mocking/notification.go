// Package mocking demonstrates how to use interfaces and mocks for testing.
// This is essential when your code depends on external services or components.
package mocking

import "fmt"

// EmailSender is an interface for sending emails.
// By using an interface, we can easily swap the real implementation
// with a mock during testing.
type EmailSender interface {
	Send(to, subject, body string) error
}

// NotificationService sends notifications to users via email.
// It depends on an EmailSender, which we can mock in tests.
type NotificationService struct {
	emailSender EmailSender
}

// NewNotificationService creates a new notification service.
func NewNotificationService(emailSender EmailSender) *NotificationService {
	return &NotificationService{
		emailSender: emailSender,
	}
}

// NotifyUser sends a notification email to a user.
func (s *NotificationService) NotifyUser(email, message string) error {
	subject := "Notification"
	return s.emailSender.Send(email, subject, message)
}

// NotifyMultipleUsers sends the same notification to multiple users.
func (s *NotificationService) NotifyMultipleUsers(emails []string, message string) error {
	for _, email := range emails {
		if err := s.NotifyUser(email, message); err != nil {
			return fmt.Errorf("failed to notify %s: %w", email, err)
		}
	}
	return nil
}

// SMTPEmailSender is a real implementation of EmailSender.
// This would actually send emails in production.
type SMTPEmailSender struct {
	host string
	port int
}

// NewSMTPEmailSender creates a new SMTP email sender.
func NewSMTPEmailSender(host string, port int) *SMTPEmailSender {
	return &SMTPEmailSender{
		host: host,
		port: port,
	}
}

// Send sends an email using SMTP.
// In a real implementation, this would connect to an SMTP server.
func (s *SMTPEmailSender) Send(to, subject, body string) error {
	// In a real implementation, this would send an actual email
	// For now, we just simulate success
	fmt.Printf("Sending email to %s: %s\n", to, subject)
	return nil
}
