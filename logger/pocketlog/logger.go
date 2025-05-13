package pocketlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

type Logger struct {
	threshold Level
	output    io.Writer
	maxLength uint
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	logger := &Logger{
		threshold: threshold,
		output:    os.Stdout,
		maxLength: 0,
	}

	for _, configFunc := range opts {
		configFunc(logger)
	}

	return logger
}

// Debugf formats and prints a message if the log level is DEBUG or higher.
func (l *Logger) Debugf(format string, args ...any) {
	l.Logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is INFO or higher.
func (l *Logger) Infof(format string, args ...any) {
	l.Logf(LevelInfo, format, args...)
}

// Errorf formats and prints a message if the log level is ERROR or higher.
func (l *Logger) Errorf(format string, args ...any) {
	l.Logf(LevelError, format, args...)
}

// Logf formats and prints a message if the log level is high enough.
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}

	l.logf(lvl, format, args...)
}

func (l *Logger) logf(lvl Level, format string, args ...any) {
	contents := fmt.Sprintf(format, args...)
	contents = capString(contents, l.maxLength)

	msg := message{
		Message: contents,
		Level:   lvl.String(),
	}

	formattedMessage, err := json.Marshal(msg)
	if err != nil {
		_, _ = fmt.Fprintf(l.output, "unable to format message for %v\n", contents)
		return
	}

	_, _ = fmt.Fprintln(l.output, string(formattedMessage))

}

func capString(s string, maxLength uint) string {
	if maxLength == 0 || uint(utf8.RuneCountInString(s)) < maxLength {
		return s
	}

	runes := []rune(s)
	cappedRunes := runes[:maxLength]

	return string(cappedRunes) + "[TRIMMED]"
}

type message struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}
