package services

import (
    "project/db"
    "project/models"
    "fmt"
)

// Service to check user existence and add the user
func CreateUser(user models.User) error {
    if db.UserExists(user.Email) {
        return fmt.Errorf("user with email %s already exists", user.Email)
    }
    return db.AddUser(user)
}
