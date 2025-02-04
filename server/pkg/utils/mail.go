package utils

import "fmt"

// BuildEmailMessage constructs an email message with the provided subject, recipient, and body.
func BuildEmailMessage(to, subject, body string) []byte {
	return []byte(fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		to,
		subject,
		body,
	))
}
