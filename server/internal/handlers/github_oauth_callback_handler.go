package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth/gothic"
	"github.com/thecodingmontana/go-community/internal/database/models"
	"github.com/thecodingmontana/go-community/pkg/types"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func (api *APIConfig) GithubOauthCallback(res http.ResponseWriter, req *http.Request) {
	req = req.WithContext(context.WithValue(req.Context(), "provider", "github"))

	auth_user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		log.Print(err)
		utils.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	// check if email is already used
	// check exisitng auth_user
	userExists, userExistsErr := api.DB.FindUserByEmail(req.Context(), auth_user.Email)

	if userExistsErr != nil {
		if errors.Is(userExistsErr, sql.ErrNoRows) {
			// auth_user doesn't exists create new auth_user
			// Create new auth_user
			newUser, createErr := api.DB.CreateUser(req.Context(), models.CreateUserParams{
				ID:            uuid.New().String(),
				Email:         auth_user.Email,
				Username:      auth_user.NickName,
				EmailVerified: true,
				Avatar: pgtype.Text{
					String: auth_user.AvatarURL,
					Valid:  true,
				},
			})
			if createErr != nil {
				utils.RespondWithError(res, http.StatusInternalServerError, "Failed to create auth_user!")
				return
			}
			// add oauth account
			_, oautAccountErr := api.DB.AddOauthAccount(req.Context(), models.AddOauthAccountParams{
				ID:             uuid.New().String(),
				UserID:         newUser.ID,
				Provider:       "github",
				ProviderUserID: auth_user.UserID,
				AccessToken:    auth_user.AccessToken,
				RefreshToken: pgtype.Text{
					String: auth_user.RefreshToken,
					Valid:  true,
				},
				ExpiresAt: pgtype.Timestamptz{
					Time:  auth_user.ExpiresAt,
					Valid: true,
				},
			})

			if oautAccountErr != nil {
				utils.RespondWithError(res, http.StatusInternalServerError, "Failed to add oauth account!")
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
			return
		} else {
			utils.RespondWithError(res, http.StatusInternalServerError, "Failed to to find auth_user with email")
			return
		}
	}

	// Generate JWT token
	claims := map[string]interface{}{
		"user_id":    userExists.ID,
		"username":   userExists.Username,
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
