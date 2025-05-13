package pocketlog

import "fmt"

type Logger struct {
	threshold Level
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and prints a message if the log level is DEBUG or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	fmt.Printf(format+"\n", args...)
}

// Infof formats and prints a message if the log level is INFO or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	fmt.Printf(format+"\n", args...)
}

// Errorf formats and prints a message if the log level is ERROR or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	fmt.Printf(format+"\n", args...)
}
