package sample_go_module

import (
	"fmt"
	"log"
)

// Log Logs a message, and returns it as a string
func Log(message string) string {
	log.Println(message)

	return message
}

// LogError Logs an error message
func LogError(message string) {
	Log(fmt.Sprintf("Error: %s", message))
}

// LogWarning Logs a warning message
func LogWarning(message string) {
	Log(fmt.Sprintf("Warning: %s", message))
}

// LogFatal Panics and logs the message twice
func LogFatal(message string) {
	panic(Log(fmt.Sprintf("Fatal: %s", message)))
}
