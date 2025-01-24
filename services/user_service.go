package services

import (
	"errors"
	"techtalk_logging/db"
	"techtalk_logging/logger"
	"techtalk_logging/models"
)

// Service to check user existence and add the user
func CheckUser(transactionID string, userName string, logger *logger.Logger) error {
	if err := checkUser(userName); err != nil {
		logger.Warn(transactionID, "check user failed with primary source for user name %s; error %v", userName, err)
	}
	return checkUserFallback(userName)
}

func checkUser(userName string) error {
	return errors.New("some error")
}

func checkUserFallback(userName string) error {
	return db.UserExists(userName)
}

func AddUser(user models.User) error {
	return db.AddUser(user)
}
