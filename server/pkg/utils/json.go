package utils

import (
	"encoding/json"
	"net/http"

	"github.com/thecodingmontana/go-community/pkg/types"
)

func RespondWithJSON(res http.ResponseWriter, status int, payload interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(payload)
}

func RespondWithError(res http.ResponseWriter, status int, message string) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)

	payload := types.APIResponse{
		StatusCode:    status,
		StatusMessage: message,
	}
	json.NewEncoder(res).Encode(payload)
}
