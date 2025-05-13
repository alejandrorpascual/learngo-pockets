package pocketlog_test

import (
	"learngo-pockets/logger/pocketlog"
	"os"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
	//Output: Hello, world

}
