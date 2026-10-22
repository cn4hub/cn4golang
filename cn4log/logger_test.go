// Package cn4log_test is used for testing.
package cn4log_test

import "github.com/cn4hub/cn4golang/cn4log"

func ExampleLogger_Debugf() {
	debugLogger := cn4log.New(cn4log.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
