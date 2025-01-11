package handlers

import (
    "encoding/json"
    "net/http"
    "project/services"
    "project/models"
)

// Handler to create a user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := services.CreateUser(user); err != nil {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
