package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/thecodingmontana/go-community/internal/database/models"
	"github.com/thecodingmontana/go-community/internal/handlers"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func RegisterRoutes(router chi.Router, queries *models.Queries, tokenAuth *jwtauth.JWTAuth) {
	apiService := handlers.NewAPIConfig(queries, tokenAuth)

	router.Route("/v1/api", func(apiRoute chi.Router) {
		apiRoute.Get("/healthz", handlers.HealthzHandler)

		apiRoute.Route("/auth", func(authroute chi.Router) {
			authroute.Post("/signin", apiService.SigninAuthHandler)
			authroute.Post("/signin/send-unique-code", apiService.SigninAuthSendVerificationCode)
			authroute.Post("/signup", apiService.SignupAuthHandler)
			authroute.Post("/signup/send-unique-code", apiService.SignupAuthSendVerificationCode)
			authroute.Get("/signin/oauth/{provider}", apiService.OauthStartAuth)
			authroute.Get("/signin/oauth/github/callback", apiService.GithubOauthCallback)
			authroute.Get("/signin/oauth/google/callback", apiService.GoogleOauthCallback)

			// protected user auth routes (JWT required)
			authroute.Group(func(protected chi.Router) {
				protected.Use(jwtauth.Verifier(tokenAuth))
				protected.Use(jwtauth.Authenticator(tokenAuth))

				protected.Get("/user", apiService.UserHandler)
			})
		})

		// protected ws routes (JWT required)
		apiRoute.Route("/ws", func(wsRoute chi.Router) {
			hub := utils.NewHub()
			go hub.Run()

			// wsRoute.Use(jwtauth.Verifier(tokenAuth))
			// wsRoute.Use(jwtauth.Authenticator(tokenAuth))

			wsRoute.Get("/chat", func(res http.ResponseWriter, req *http.Request) {
				handlers.ServeWSChatHandler(res, req, hub, queries)
			})
		})
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
