package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func GenerateUsernameFromEmail(email string) string {
	// Remove spaces and convert to lowercase
	email = strings.ToLower(strings.TrimSpace(email))

	// Get the part before @ symbol
	username := strings.Split(email, "@")[0]

	// Remove special characters and numbers
	cleaned := ""
	for _, char := range username {
		if unicode.IsLetter(char) {
			cleaned += string(char)
		}
	}

	// Handle edge cases
	if len(cleaned) == 0 {
		return "user"
	}

	// Capitalize first letter
	if len(cleaned) > 0 {
		cleaned = strings.ToUpper(cleaned[:1]) + cleaned[1:]
	}

	return cleaned
}

// GenerateUniqueUsername ensures the username is unique by adding a number if needed
func GenerateUniqueUsername(email string, existingUsernames map[string]bool) string {
	baseUsername := GenerateUsernameFromEmail(email)
	username := baseUsername
	counter := 1

	// Keep adding numbers until we find a unique username
	for existingUsernames[username] {
		username = baseUsername + strings.TrimSpace(fmt.Sprint(counter))
		counter++
	}

	return username
}
