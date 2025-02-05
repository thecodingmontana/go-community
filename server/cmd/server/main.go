package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/thecodingmontana/go-community/internal/database/models"
	"github.com/thecodingmontana/go-community/internal/routes"
	"github.com/thecodingmontana/go-community/pkg/database"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file, probably missing!")
	}

	JWT_SECRET, jwtSecretFound := os.LookupEnv("JWT_SECRET")
	if !jwtSecretFound {
		log.Fatalf("Missing JWT_SECRET env variable!")
	}
	tokenAuth = jwtauth.New("HS256", []byte(JWT_SECRET), nil)
}

func main() {
	PORT, isPortEnv := os.LookupEnv("PORT")
	DATABASE_URL, isDBURLEnv := os.LookupEnv("DATABASE_URL")
	OAUTH_GOOGLE_CLIENT_ID, googleClientIDFound := os.LookupEnv("OAUTH_GOOGLE_CLIENT_ID")
	OAUTH_GOOGLE_CLIENT_SECRET, googleClientSecretFound := os.LookupEnv("OAUTH_GOOGLE_CLIENT_SECRET")
	OAUTH_GOOGLE_REDIRECT_URI, googleRedirectURIFound := os.LookupEnv("OAUTH_GOOGLE_REDIRECT_URI")
	OAUTH_GITHUB_CLIENT_ID, githubClientIDFound := os.LookupEnv("OAUTH_GITHUB_CLIENT_ID")
	OAUTH_GITHUB_CLIENT_SECRET, githubClientIDFound := os.LookupEnv("OAUTH_GITHUB_CLIENT_SECRET")
	OAUTH_GITHUB_REDIRECT_URI, githubRedirectURIFound := os.LookupEnv("OAUTH_GITHUB_REDIRECT_URI")
	SESSION_SECRET, isSessionSecretFound := os.LookupEnv("SESSION_SECRET")
	ALLOWED_ORIGINS, isAllowedOrigins := os.LookupEnv("ALLOWED_ORIGINS")

	if !isPortEnv {
		log.Fatalf("Missing PORT env variable!")
	}

	if !isDBURLEnv {
		log.Fatalf("Missing DATABASE_URL env variable!")
	}

	if !googleClientIDFound {
		log.Fatalf("Missing OAUTH_GOOGLE_CLIENT_ID env variable!")
	}

	if !googleClientSecretFound {
		log.Fatalf("Missing OAUTH_GOOGLE_CLIENT_SECRET env variable!")
	}

	if !googleRedirectURIFound {
		log.Fatalf("Missing OAUTH_GOOGLE_REDIRECT_URI env variable!")
	}

	if !githubClientIDFound {
		log.Fatalf("Missing OAUTH_GITHUB_CLIENT_ID env variable!")
	}

	if !githubClientIDFound {
		log.Fatalf("Missing OAUTH_GITHUB_CLIENT_SECRET env variable!")
	}

	if !githubRedirectURIFound {
		log.Fatalf("Missing OAUTH_GITHUB_REDIRECT_URI env variable!")
	}

	if !isSessionSecretFound {
		log.Fatalf("Missing SESSION_SECRET env variable!")
	}

	if !isAllowedOrigins {
		log.Fatalf("Missing ALLOWED_ORIGINS env variable!")
	}

	originsList := strings.Split(ALLOWED_ORIGINS, ",")

	// Database connection
	conn := database.ConnectDB(DATABASE_URL)
	queries := models.New(conn)

	router := chi.NewRouter()

	maxAge := 86400 * 30 // 30 days
	isProd := true       // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(SESSION_SECRET))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	// Initialize Goth providers
	goth.UseProviders(
		google.New(
			OAUTH_GOOGLE_CLIENT_ID,
			OAUTH_GOOGLE_CLIENT_SECRET,
			OAUTH_GOOGLE_REDIRECT_URI,
			"email",
			"profile",
		),
		github.New(
			OAUTH_GITHUB_CLIENT_ID,
			OAUTH_GITHUB_CLIENT_SECRET,
			OAUTH_GITHUB_REDIRECT_URI,
			"user:email",
			"user",
		),
	)

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   originsList,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
		AllowCredentials: true,
	}))

	// Register routes and handlers
	routes.RegisterRoutes(router, queries, tokenAuth)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("🚀 Server started at http://localhost:%s", PORT)

	if serverErr := server.ListenAndServe(); serverErr != nil {
		log.Fatalf("Failed to start the server: %v", serverErr)
	}
}
