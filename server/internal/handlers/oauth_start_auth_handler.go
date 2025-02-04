package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func (api *APIConfig) OauthStartAuth(res http.ResponseWriter, req *http.Request) {
	provider := chi.URLParam(req, "provider")

	if provider == "" {
		utils.RespondWithError(res, http.StatusBadRequest, "Provider not specified")
		return
	}
	req = req.WithContext(context.WithValue(req.Context(), "provider", provider))

	gothic.BeginAuthHandler(res, req)
}
