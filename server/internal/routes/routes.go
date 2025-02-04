package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thecodingmontana/go-community/internal/database/models"
	"github.com/thecodingmontana/go-community/internal/handlers"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func RegisterRoutes(router chi.Router, queries *models.Queries) {
	router.Route("/v1/api", func(apiRoute chi.Router) {
		apiRoute.Get("/healthz", handlers.HealthzHandler)
	})

	// 404 route
	router.NotFound(func(res http.ResponseWriter, req *http.Request) {
		utils.RespondWithError(res, http.StatusNotFound, "Route does not exists!")
	})

	// Method not allowed
	router.MethodNotAllowed(func(res http.ResponseWriter, req *http.Request) {
		utils.RespondWithError(res, http.StatusMethodNotAllowed, "Method is invalid!")
	})
}
