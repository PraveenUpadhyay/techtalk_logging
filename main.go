package main

import (
	"log"
	"net/http"

	"techtalk_logging/logger"

	"techtalk_logging/handlers"
)

func main() {
	// Create a new logger instance
	logger := logger.NewLogger()
	// Pass the logger to the handler
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUserHandler(w, r, logger)
	})
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
