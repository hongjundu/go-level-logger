package main

import (
	"github.com/du-hj/go-level-logger"
)

func main() {
	logger.InitLogger(logger.LogLevelDebug)

	logger.Debugf("debug message")
	logger.Infof("info message")
	logger.Noticef("notice message")
	logger.Warnf("warn message")
	logger.Errorf("error message")
}
