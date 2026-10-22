// Package cn4log provides a logger library.
package cn4log

// Logger is used to log information.
type Logger struct {
	threshold Level
}

// Debugf formats and prints a message when the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
}

// Infof formats and prints a message when the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
}

// Warnf formats and prints a message when the log level is warn or higher.
func (l *Logger) Warnf(format string, args ...any) {
}

// Errorf formats and prints a message when the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
}

// New creates a logger with the threshold level specified.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
