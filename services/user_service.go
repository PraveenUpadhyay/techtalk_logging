package services

import (
	"fmt"
	"techtalk_logging/db"
	"techtalk_logging/logger"
	"techtalk_logging/models"
)

// Service to check user existence and add the user
func CheckUser(transactionID string, userName string, logger *logger.Logger) error {
	if err := checkUser(userName); err != nil {
		logger.Error(transactionID, "check user failed with primary source for user name %s; error %v", userName, err)
	}
	if err := checkUserFallback(userName); err != nil {
		logger.Error(transactionID, "check user failed with seconday source as well for user name %s; error %v", userName, err)
		return err
	}
	return nil
}

func checkUser(userName string) error {
	return fmt.Errorf("could not fetch from primary source for user name %s", userName)
}

func checkUserFallback(userName string) error {
	if err := db.UserExists(userName); err != nil {
		return fmt.Errorf("user already exists for user name %s; error %v", userName, err)
	}
	return nil
}

func AddUser(user models.User) error {
	return db.AddUser(user)
}
