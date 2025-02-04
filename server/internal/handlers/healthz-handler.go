package handlers

import (
	"net/http"

	"github.com/thecodingmontana/go-community/pkg/types"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func HealthzHandler(res http.ResponseWriter, req *http.Request) {
	utils.RespondWithJSON(res, http.StatusOK, types.APIResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "Everything is working fine 🤩!",
	})
}
