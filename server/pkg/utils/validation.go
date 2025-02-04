package utils

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func HandleValidation(res http.ResponseWriter, body interface{}) (bool, string) {
	// Create a new validator instance
	validate := validator.New()

	// Create a title caser using the default language (English)
	caser := cases.Title(language.English)

	// Validate the body
	if err := validate.Struct(body); err != nil {
		// Check for validator.ValidationErrors type
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Create a user-friendly error message
			var errorMessages []string
			for _, err := range validationErrors {
				// Format the error message for each failed validation
				errorMessages = append(errorMessages, caser.String(err.Field())+" is required")
			}
			// Combine the error messages into one string
			errorMessage := strings.Join(errorMessages, ", ")
			return false, errorMessage
		}
		// Handle other types of errors (if any)
		return false, "An unknown error occurred"
	}
	return true, ""
}

// Helper function to validate email format
func IsValidEmail(email string) bool {
	// Simple regex for email validation
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
