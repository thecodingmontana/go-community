package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/thecodingmontana/go-community/internal/routes"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file, probably missing!")
	}

	PORT, isPortEnv := os.LookupEnv("PORT")
	DATABASE_URL, isDBURLEnv := os.LookupEnv("DATABASE_URL")

	if !isPortEnv {
		log.Fatalf("Missing PORT env variable!")
	}

	if !isDBURLEnv {
		log.Fatalf("Missing DATABASE_URL env variable!")
	}

	log.Print(DATABASE_URL)

	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Register routes and handlers
	routes.RegisterRoutes(router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("🚀 Server started at http://localhost:%s", PORT)

	if serverErr := server.ListenAndServe(); serverErr != nil {
		log.Fatalf("😒 Failed to start the server: %v", serverErr)
	}
}
