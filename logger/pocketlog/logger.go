package pocketlog

type Logger struct {
}

// Debugf formats and prints a message if the log level is DEBUG or higher
func (l *Logger) Debugf(format string, args ...any) {
	// implement
}

// Infof formats and prints a message if the log level is INFO or higher
func (l *Logger) Infof(format string, args ...any) {
	// implement
}

// Errorf formats and prints a message if the log level is ERROR or higher
func (l *Logger) Errorf(format string, args ...any) {
	// implement
}
