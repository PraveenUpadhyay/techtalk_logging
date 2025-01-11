package db

import "project/models"

// Dummy database to hold users
var users = make(map[string]models.User)

// Check if a user already exists in the database
func UserExists(email string) bool {
    _, exists := users[email]
    return exists
}

// Add a user to the database
func AddUser(user models.User) error {
    if UserExists(user.Email) {
        return fmt.Errorf("user already exists")
    }
    users[user.Email] = user
    return nil
}
