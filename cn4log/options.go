// Package cn4log provides a logger library.
package cn4log

import (
	"io"
)

// Option defines a functional option to our logger
type Option func(*Logger)

// WithOutput sets the output for the logger
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}

// WithName sets the name for the logger
func WithName(name string) Option {
	return func(l *Logger) {
		l.name = name
	}
}

func NoLvl() Option {
	return func(l *Logger) {
		l.printLevel = false
	}
}

func NoTime() Option {
	return func(l *Logger) {
		l.printTime = false
	}
}
