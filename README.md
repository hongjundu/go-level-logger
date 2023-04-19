# go-level-logger

## Features

* Dynamically config log level
* Output debug/info/notice messages to stdout, and output warn/error mssages to stdout/stderr 
* Customize log output

## Examples

### Basic

    package main

    import (
        "github.com/hongjundu/go-level-logger"
    )

    func main() {
        logger.InitLogger(logger.LogLevelDebug)

        logger.Debugf("debug message")
        logger.Infof("info message")
        logger.Noticef("notice message")
        logger.Warnf("warn message")
        logger.Errorf("error message")
    }

### Customize log output

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

### Customize rotate log output (Rotate log  at 00:00:00 per day)

    package main

    import (
        "github.com/hongjundu/go-level-logger"
    )

    func main() {
        logger.InitRotateLogger(logger.LogLevelDebug, "example", "/tmp", 100, 3, 365, logger.RotateStrategyPerDay)

        logger.Debugf("debug message")
        logger.Infof("info message")
        logger.Noticef("notice message")
        logger.Warnf("warn message")
        logger.Errorf("error message")
    }

### Customize rotate log output (Rotate log  per hour)

    package main

    import (
        "github.com/hongjundu/go-level-logger"
    )

    func main() {
        logger.InitRotateLogger(logger.LogLevelDebug, "example", "/tmp", 100, 3, 365, logger.RotateStrategyPerHour)

        logger.Debugf("debug message")
        logger.Infof("info message")
        logger.Noticef("notice message")
        logger.Warnf("warn message")
        logger.Errorf("error message")
    }