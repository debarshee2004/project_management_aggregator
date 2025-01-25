package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	// Get the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	// Create a new router
	router := chi.NewRouter()

	// Set up CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create a sub-router for version 1 of the API
	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness) // Health check endpoint
	v1Router.Get("/err", handlerErr)          // Error endpoint for testing

	// Mount the version 1 router to the main router
	router.Mount("/v1", v1Router)

	// Create and start the HTTP server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// Listen and serve on the specified port
	error := server.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}

	log.Println("Server started on port " + port)
}
