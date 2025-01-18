package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"techtalk_logging/logger"

	"techtalk_logging/models"
	"techtalk_logging/services"

	"github.com/google/uuid"
)

// Handler to create a user
func CreateUserHandler(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	// Generate a new UUID
	transactionID := uuid.New().String()

	var user models.User
	payload, _ := io.ReadAll(r.Body)
	logger.Info(transactionID, "received request to create user for payload %s", payload)
	if err := json.Unmarshal(payload, &user); err != nil {
		logger.Error(transactionID, "failed to process request %s; error %v", payload, err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.CheckUser(transactionID, user.Name, logger); err != nil {
		logger.Error(transactionID, "failed to create due to conflict for payload %v; error %v", user, err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	if err := services.AddUser(user); err != nil {
		logger.Error(transactionID, "failed to create user for payload %v; error %v", user, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info(transactionID, "user created successfully for payload %s", payload)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
