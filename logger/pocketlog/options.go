package pocketlog

import "io"

// Option defines a functional option to our logger.
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output of
// logs.
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}

// WithMaxLength returns a configuration function that sets the maximum
// length of the log messages.
func WithMaxLength(maxLength uint) Option {
	return func(l *Logger) {
		l.maxLength = maxLength
	}
}
