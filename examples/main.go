package main

import (
	"github.com/hongjundu/go-level-logger"
)

func main() {
	logger.Init(logger.LogLevelDebug, "example", "/tmp", 100, 3, 365)

	logger.Debugf("debug message")
	logger.Infof("info message")
	logger.Noticef("notice message")
	logger.Warnf("warn message")
	logger.Errorf("error message")
}
