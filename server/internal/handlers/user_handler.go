package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/thecodingmontana/go-community/pkg/types"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func (api *APIConfig) UserHandler(res http.ResponseWriter, req *http.Request) {
	_, claims, _ := jwtauth.FromContext(req.Context())

	userId, ok := claims["user_id"].(string)
	if !ok {
		utils.RespondWithError(res, http.StatusUnauthorized, "Invalid user ID type in claims")
		return
	}

	user, err := api.DB.FindUserByID(req.Context(), userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(res, http.StatusBadRequest, "User not found!")
			return
		}
		utils.RespondWithError(res, http.StatusInternalServerError, "Failed to find user by ID!")
		return
	}

	data := types.User{
		ID:            user.ID,
		Email:         user.Email,
		Username:      user.Username,
		EmailVerified: user.EmailVerified,
		Avatar:        user.Avatar.String,
	}

	log.Println(data)

	utils.RespondWithJSON(res, http.StatusOK, types.APIResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "User found!",
		Data:          data,
	})
}
