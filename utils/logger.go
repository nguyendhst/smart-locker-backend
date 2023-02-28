package utils

import (
	"fmt"
	"log"
	"os"
)

// Logger is a wrapper around the standard log package.
type Logger struct {
	module string
	*log.Logger
	isStdout bool
	logFile  *os.File
}

// NewLogger creates a new logger.
func NewLogger(module string) *Logger {
	return &Logger{
		Logger:   log.New(os.Stdout, "", log.LstdFlags),
		isStdout: true,
		module:   module,
	}
}

// Info logs an info message.
func (l *Logger) Info(msg string) error {
	// If the logger is writing to stdout, then we need to add a newline.
	// INFO in blue.
	if l.isStdout {
		l.Printf("\x1b[34m%s-INFO\x1b[0m: %s \x1b[0m ", l.module, msg)
	} else {
		// write to log file
		if _, err := l.logFile.WriteString(fmt.Sprintf("%s-INFO: %s \x1b[0m ", l.module, msg)); err != nil {
			return err
		}
	}
	return nil
}
