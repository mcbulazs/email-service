package logging

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"mcbulazs/email-service/internal/config"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func Init() {
	logFile := config.AppConfig.LogFile

	// Ensure log directory exists
	if err := os.MkdirAll(filepath.Dir(logFile), os.ModePerm); err != nil {
		log.Fatalf("Could not create log directory: %v", err)
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Could not open log file: %v", err)
	}

	multiInfo := io.MultiWriter(file, os.Stdout)
	multiError := io.MultiWriter(file, os.Stderr)

	InfoLogger = log.New(multiInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multiError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
