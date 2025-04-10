package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Cloud Run injects the PORT environment variable.
	// Default to 8080 for local development.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Define a simple handler function.
	handler := func(w http.ResponseWriter, r *http.Request) {
		// You can access environment variables provided during deployment or locally.
		// Example: GREETING environment variable
		greeting := os.Getenv("GREETING")
		if greeting == "" {
			greeting = "Hello"
		}

		// Get the target from the query parameter, default to "World"
		target := r.URL.Query().Get("target")
		if target == "" {
			target = "World"
		}

		log.Printf("Serving request for path: %s", r.URL.Path) // Log request path
		fmt.Fprintf(w, "%s, %s!\n", greeting, target)
	}

	// Register the handler function for the root path "/".
	// All incoming requests will be handled by this function.
	http.HandleFunc("/", handler)

	// Start the HTTP server.
	log.Printf("Listening on port %s", port)
	// Listen on all network interfaces (":<port>") which is required for Cloud Run.
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
