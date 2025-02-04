package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thecodingmontana/go-community/internal/database/models"
	"github.com/thecodingmontana/go-community/pkg/mail"
	"github.com/thecodingmontana/go-community/pkg/mail/templates"
	"github.com/thecodingmontana/go-community/pkg/types"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func (api APIConfig) SignupAuthSendVerificationCode(res http.ResponseWriter, req *http.Request) {
	var body struct {
		Email string `validate:"required,email"`
	}

	// Decode the request body
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		utils.RespondWithError(res, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer req.Body.Close()

	// Validate the body
	isValid, validationErr := utils.HandleValidation(res, body)
	if !isValid {
		utils.RespondWithError(res, http.StatusBadRequest, validationErr)
		return
	}

	// Check if the user exists
	user, findErr := api.DB.FindUserByEmail(req.Context(), body.Email)
	if findErr == nil {
		// User exists, return error
		utils.RespondWithError(res, http.StatusBadRequest, fmt.Sprintf("Email %s already exists!", user.Email))
		return
	} else if !errors.Is(findErr, sql.ErrNoRows) {
		// Unexpected database error
		log.Printf("Unexpected database error while finding user by email: %v", findErr)
		utils.RespondWithError(res, http.StatusInternalServerError, "Database error while finding user by email")
		return
	}

	// Generate new code
	code := utils.GenerateRandomString(8)
	now := time.Now()
	expiresAt := now.Add(10 * time.Minute)

	// Check if unique code exists
	_, codeExistsError := api.DB.FindUserUniqueCode(req.Context(), body.Email)
	if codeExistsError != nil && !errors.Is(codeExistsError, sql.ErrNoRows) {
		utils.RespondWithError(res, http.StatusInternalServerError, "Database error while checking unique code")
		return
	}

	if errors.Is(codeExistsError, sql.ErrNoRows) {
		// No code exists yet, so create a new one
		result, createErr := api.DB.GenerateUniqueCode(req.Context(), models.GenerateUniqueCodeParams{
			ID:    uuid.New().String(),
			Email: body.Email,
			Code:  code,
			ExpiresAt: pgtype.Timestamptz{
				Time:  expiresAt,
				Valid: true,
			},
		})
		if createErr != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, "Failed to generate unique code")
			return
		}

		// Set up the message
		message := []byte(
			fmt.Sprintf(
				"To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
				body.Email,
				fmt.Sprintf("Your unique use-odama sign up code is %s", result.Code),
				templates.UniqueCodeTemplate(),
			),
		)

		// Send email
		mail.AppSendMail([]string{body.Email}, types.MailBody{
			Message: message,
		})
	} else {
		// Update existing code
		updatedCode, updateErr := api.DB.UpdateUniqueCode(req.Context(), models.UpdateUniqueCodeParams{
			Code: code,
			ExpiresAt: pgtype.Timestamptz{
				Time:  expiresAt,
				Valid: true,
			},
			UpdatedAt: pgtype.Timestamp{
				Time:  now,
				Valid: true,
			},
			Email: body.Email,
		})

		if updateErr != nil {
			log.Println(updateErr)
			utils.RespondWithError(res, http.StatusInternalServerError, "Failed to update unique code")
			return
		}

		// Set up the message
		message := []byte(
			fmt.Sprintf(
				"To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
				body.Email,
				fmt.Sprintf("Your unique use-odama verification code is %s", updatedCode),
				templates.UniqueCodeTemplate(),
			),
		)

		// Send email
		mail.AppSendMail([]string{body.Email}, types.MailBody{
			Message: message,
		})
	}

	// Send success response
	utils.RespondWithJSON(res, http.StatusOK, types.APIResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "Check your email for the verification code!",
	})
}
