// Package cn4log provides a logger library.
package cn4log

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold  Level
	name       string
	output     io.Writer
	printTime  bool
	printLevel bool
}

// logf formats and prints a message when the given level is equal or higher to the one specified in the logger.
func (l *Logger) logf(level Level, format string, args ...any) {
	if l.threshold > level {
		return
	}
	out := l.output
	if out == nil {
		out = os.Stdout
	}
	fmt.Fprintf(out, format+"\n", args...)
}

// logf formats and prints a message when the given level is equal or higher to the one specified in the logger.
func (l *Logger) log(level Level, msg string) {
	if l.threshold > level {
		return
	}
	out := l.output
	if out == nil {
		out = os.Stdout
	}
	fmt.Fprintln(out, msg)
}

// Debugf formats and prints a message when the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	l.logf(LevelDebug, format, args...)
}

// Debug prints a message when the log level is debug or higher.
func (l *Logger) Debug(msg string) {
	l.log(LevelDebug, msg)
}

// Infof formats and prints a message when the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	l.logf(LevelInfo, format, args...)
}

// Info prints a message when the log level is debug or higher.
func (l *Logger) Info(msg string) {
	l.log(LevelInfo, msg)
}

// Warnf formats and prints a message when the log level is warn or higher.
func (l *Logger) Warnf(format string, args ...any) {
	l.logf(LevelWarn, format, args...)
}

// Warn prints a message when the log level is debug or higher.
func (l *Logger) Warn(msg string) {
	l.log(LevelWarn, msg)
}

// Errorf formats and prints a message when the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	l.logf(LevelError, format, args...)
}

// Error prints a message when the log level is debug or higher.
func (l *Logger) Error(msg string) {
	l.log(LevelError, msg)
}

// New creates a logger with the threshold level specified.
func New(threshold Level, opts ...Option) *Logger {
	l := &Logger{
		threshold:  threshold,
		output:     os.Stdout,
		printTime:  true,
		printLevel: true,
	}
	for _, configFunc := range opts {
		configFunc(l)
	}
	return l
}
