package go_logger

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

// New creates a new logger instance
func New(nodeConfig string) *Logger {
	initialize(nodeConfig)
	return &Logger{
		logger: log.New(os.Stdout, nodeConfig, log.LstdFlags),
	}
}

// Info logs an informational message
func (l *Logger) Info(msg string) {
	l.logger.Printf("[INFO] %s", msg)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.logger.Printf("[ERROR] %s", msg)
}
