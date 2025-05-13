package pocketlog_test

import (
	"learngo-pockets/logger/pocketlog"
	"os"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout))
	debugLogger.Debugf("Hello, %s", "world")
	//Output: [DEBUG] Hello, world

}

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	var (
		testDebugMessage = "[DEBUG] " + debugMessage + "\n"
		testInfoMessage  = "[INFO] " + infoMessage + "\n"
		testErrorMessage = "[ERROR] " + errorMessage + "\n"
	)

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: testDebugMessage + testInfoMessage + testErrorMessage,
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: testInfoMessage + testErrorMessage,
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: testErrorMessage,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testLogger.Debugf(debugMessage)
			testLogger.Infof(infoMessage)
			testLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

// testWriter is a struct that implements io.Writer.
// We use it to validate that we can writet o a specific output.
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
