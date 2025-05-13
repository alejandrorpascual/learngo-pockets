package pocketlog

import (
	"fmt"
	"io"
)

type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// Debugf formats and prints a message if the log level is DEBUG or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf(format, args...)
}

// Infof formats and prints a message if the log level is INFO or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf(format, args...)
}

// Errorf formats and prints a message if the log level is ERROR or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.logf(format, args...)
}

func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
