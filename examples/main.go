package main

import (
	"github.com/hongjundu/go-level-logger"
	"github.com/petermattis/goid"
)

func main() {
	logger.Init(logger.LogLevelDebug, "example", "/tmp", 100, 3, 365)

	logger.Debugf("debug message")
	logger.Infof("info message")
	logger.Noticef("notice message")
	logger.Warnf("warn message")
	logger.Errorf("error message")

	testTrace()

	logger.Debugf("no trace")
}

func testTrace() {
	logger.StoreTraceId(goid.Get(), "63b0359fee4eebf5")
	defer logger.ClearTraceId(goid.Get())

	logger.Debugf("debug trace message")
	logger.Infof("info trace message")
	logger.Noticef("notice trace message")
	logger.Warnf("warn trace message")
	logger.Errorf("error trace message")
}
