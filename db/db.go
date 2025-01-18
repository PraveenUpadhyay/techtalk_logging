package db

import (
	"techtalk_logging/common"
	"techtalk_logging/models"
)

// Dummy database to hold users
var users = make(map[string]models.User)

// Check if a user already exists in the database
func UserExists(userName string) error {
	if _, exists := users[userName]; exists {
		return common.ErrExists
	}
	return nil
}

// Add a user to the database
func AddUser(user models.User) error {
	users[user.Name] = user
	return nil
}
