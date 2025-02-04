package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/thecodingmontana/go-community/pkg/types"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func (api APIConfig) SigninAuthHandler(res http.ResponseWriter, req *http.Request) {
	var body struct {
		Email string `validate:"required,email"`
		Code  string `validate:"required,len=8"`
	}

	// Decode the request body
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		utils.RespondWithError(res, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer req.Body.Close()

	// Validate the request body
	if isValid, validationErr := utils.HandleValidation(res, body); !isValid {
		utils.RespondWithError(res, http.StatusBadRequest, validationErr)
		return
	}

	// Fetch the user by email
	user, err := api.DB.FindUserByEmail(req.Context(), body.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(res, http.StatusNotFound, fmt.Sprintf("Email %s doesn't exist!", body.Email))
			return
		}
		log.Printf("Error finding user: %v", err)
		utils.RespondWithError(res, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	// Fetch the unique code for the user
	uniqueCode, err := api.DB.FindUserUniqueCode(req.Context(), body.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid unique code!")
			return
		}
		log.Printf("Error finding unique code: %v", err)
		utils.RespondWithError(res, http.StatusInternalServerError, "Database error while retrieving unique code")
		return
	}

	// Validate the unique code
	if uniqueCode.Code != body.Code {
		utils.RespondWithError(res, http.StatusBadRequest, "Invalid unique code!")
		return
	}

	// Check if the unique code has expired
	if time.Now().After(uniqueCode.ExpiresAt.Time) {
		utils.RespondWithError(res, http.StatusBadRequest, "Unique code has expired!")
		return
	}

	// Delete the unique code request
	if err := api.DB.DeleteUniqueCodeRequest(req.Context(), body.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(res, http.StatusNotFound, fmt.Sprintf("No unique code found for email: %s", body.Email))
			return
		}
		log.Printf("Error deleting unique code: %v", err)
		utils.RespondWithError(res, http.StatusInternalServerError, "Failed to delete unique code")
		return
	}

	// Generate JWT token
	claims := map[string]interface{}{
		"user_id":    user.ID,
		"username":   user.Username,
		"expires_at": time.Now().Add(24 * time.Hour).Unix(),
	}

	_, tokenString, err := api.tokenAuth.Encode(claims)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		utils.RespondWithError(res, http.StatusInternalServerError, "Error generating token")
		return
	}

	// Respond with the token and expiration
	utils.RespondWithJSON(res, http.StatusOK, types.APIResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "Successfully signed in!",
		Data: map[string]interface{}{
			"token":      tokenString,
			"expires_at": claims["expires_at"],
		},
	})
}
