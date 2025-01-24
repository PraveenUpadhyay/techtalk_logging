package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"techtalk_logging/common"
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
	if err := json.Unmarshal(payload, &user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.CheckUser(transactionID, user.Name, logger); err != nil {
		if err == common.ErrExists {
			//4xx does not need to be logged
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		logger.Error(transactionID, "failed to check if user exist for payload %v; error %v", user, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
