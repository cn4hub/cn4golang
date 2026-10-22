// Package cn4log provides a logger library.
package cn4log

// Level represents an available logging level.
type Level byte

const (
	// LevelDebug is the lowest level and shall be used by develpers.
	LevelDebug = iota
	// LevelInfo contains general information.
	LevelInfo
	// LevelWarn is for unusual critical information that are not of level LevelError.
	LevelWarn
	// LevelError is used for tracing errors.
	LevelError
)
