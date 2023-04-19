package main

import (
	"github.com/hongjundu/go-level-logger"
	"time"
)

func main() {
	//logger.Init(logger.LogLevelDebug, "example", "/tmp", 100, 3, 365)
	//
	//logger.Debugf("debug message")
	//logger.Infof("info message")
	//logger.Noticef("notice message")
	//logger.Warnf("warn message")
	//logger.Errorf("error message")

	logger.InitRotateLogger(logger.LogLevelDebug, "example", "/tmp", 100, 3, 365, logger.RotateStrategyPerHour)

	for idx := 0; idx < 1000000; idx++ {
		time.Sleep(1 * time.Second)
		logger.Debugf("debug message")
		logger.Infof("info message")
		logger.Noticef("notice message")
		logger.Warnf("warn message")
		logger.Errorf("error message")
	}

}
