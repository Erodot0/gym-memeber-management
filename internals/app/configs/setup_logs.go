package configs

import (
	"log"
	"os"
	"path/filepath"
)

func InitLogs() {
	// Define the log directory and file
	logDir := "./logs"
	logFile := filepath.Join(logDir, "app.log")

	// Create the log directory if it does not exist
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create log directory: %v", err)
	}


	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
