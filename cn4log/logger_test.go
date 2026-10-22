// Package cn4log_test is used for testing.
package cn4log_test

import (
	"testing"

	"github.com/cn4hub/cn4golang/cn4log"
)

const (
	errorMsg = "Sei du selbst die Veränderung, die du dir wünschst für diese Welt."
	warnMsg  = "Der einzige Weg, großartige Arbeit zu leisten, besteht darin, zu lieben, was man tut."
	infoMsg  = "Es ist nicht wichtig, wie langsam man geht, solange man nicht stehen bleibt."
	debugMsg = "Life is like riding a bicycle. To keep your balance, you must keep moving."
)

func ExampleLogger_Debugf() {
	debugLogger := cn4log.New(cn4log.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}

// testWriter is a struct that implements io.Writer.
// We use it to validaate that we can write to a specific output.
type testWriter struct {
	contents string
}

// Writer implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    cn4log.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    cn4log.LevelDebug,
			expected: debugMsg + "\n" + infoMsg + "\n" + warnMsg + "\n" + errorMsg + "\n",
		},
		"info": {
			level:    cn4log.LevelInfo,
			expected: infoMsg + "\n" + warnMsg + "\n" + errorMsg + "\n",
		},
		"warn": {
			level:    cn4log.LevelWarn,
			expected: warnMsg + "\n" + errorMsg + "\n",
		},
		"error": {
			level:    cn4log.LevelError,
			expected: errorMsg + "\n",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := cn4log.New(tc.level, cn4log.WithOutput(tw))
			testedLogger.Debugf(debugMsg)
			testedLogger.Infof(infoMsg)
			testedLogger.Warnf(warnMsg)
			testedLogger.Errorf(errorMsg)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}
