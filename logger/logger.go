package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger struct that holds a logger instance
type Logger struct {
	*log.Logger
}

// NewLogger creates and returns a new Logger instance
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// INFO logs informational messages
func (l *Logger) Info(transactionID string, message string, v ...interface{}) {
	l.Logger.SetPrefix(fmt.Sprintf("%s: %s : ", "INFO", transactionID))
	l.Logger.Printf(message, v...)
}

// ERROR logs error messages
func (l *Logger) Error(transactionID, message string, v ...interface{}) {
	l.Logger.SetPrefix(fmt.Sprintf("%s: %s : ", "ERROR", transactionID))
	l.Logger.Printf(message, v...)
}

// WARN logs error messages
func (l *Logger) Warn(transactionID, message string, v ...interface{}) {
	l.Logger.SetPrefix(fmt.Sprintf("%s: %s : ", "WARN", transactionID))
	l.Logger.Printf(message, v...)
}
