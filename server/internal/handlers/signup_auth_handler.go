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
	"github.com/thecodingmontana/go-community/pkg/types"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func (api APIConfig) SignupAuthHandler(res http.ResponseWriter, req *http.Request) {
	var body struct {
		Email string `validate:"required,email"`
		Code  string `validate:"required,max=8,min=8"`
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

	username := utils.GenerateUsernameFromEmail(body.Email)

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

	// Check if unique code exists
	uniqueCodeRequest, codeExistsError := api.DB.FindUserUniqueCode(req.Context(), body.Email)
	if codeExistsError != nil {
		if errors.Is(codeExistsError, sql.ErrNoRows) {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid unique code!")
		} else {
			utils.RespondWithError(res, http.StatusInternalServerError, "Database error while checking unique code")
		}
		return
	}

	if uniqueCodeRequest.Code != body.Code {
		utils.RespondWithError(res, http.StatusBadRequest, "Invalid unique code!")
		return
	}

	// Time has expired
	if time.Now().After(uniqueCodeRequest.ExpiresAt.Time) {
		utils.RespondWithError(res, http.StatusBadRequest, "Unique code has expired!")
		return
	}

	// Delete unique code request
	deleteErr := api.DB.DeleteUniqueCodeRequest(req.Context(), body.Email)
	if deleteErr != nil {
		if errors.Is(deleteErr, sql.ErrNoRows) {
			utils.RespondWithError(res, http.StatusNotFound, fmt.Sprintf("No unique code found for email: %s", body.Email))
			return
		}
		utils.RespondWithError(res, http.StatusInternalServerError, "Failed to delete unique code")
		return
	}

	// Create new user
	newUser, createErr := api.DB.CreateUser(req.Context(), models.CreateUserParams{
		ID:            uuid.New().String(),
		Email:         body.Email,
		Username:      username,
		EmailVerified: true,
		Avatar: pgtype.Text{
			String: fmt.Sprintf("https://avatar.vercel.sh/%s", body.Email),
			Valid:  true,
		},
	})
	if createErr != nil {
		utils.RespondWithError(res, http.StatusInternalServerError, "Failed to create user!")
		return
	}

	// Create JWT claims
	claims := map[string]interface{}{
		"user_id":    newUser.ID,
		"username":   newUser.Username,
		"expires_at": time.Now().Add(24 * time.Hour).Unix(),
	}

	_, tokenString, tokenErr := api.tokenAuth.Encode(claims)
	if tokenErr != nil {
		log.Printf("Error generating token: %v", tokenErr)
		utils.RespondWithError(res, http.StatusInternalServerError, "Error generating token")
		return
	}

	data := map[string]interface{}{
		"token":      tokenString,
		"expires_at": claims["expires_at"].(int64),
	}

	// Send success response
	utils.RespondWithJSON(res, http.StatusOK, types.APIResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "Successfully signed up!.",
		Data:          data,
	})
}
