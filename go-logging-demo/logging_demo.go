package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new logger
	logger := logrus.New()

	// Set output to file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	// Set log level
	logger.SetLevel(logrus.InfoLevel)

	// Log with different levels
	logger.Info("Application started")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// Log with fields
	logger.WithFields(logrus.Fields{
		"user_id": 12345,
		"action":  "login",
	}).Info("User logged in")

	// Standard library logging
	logrus.Info("Using standard logrus")
	logrus.WithField("component", "main").Info("Component logging")
}