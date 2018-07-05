# go-level-logger

## Features

* Dynamically config log level
* Output debug/info/notice messages to stdout, and output warn/error mssages to both stdout and stderr 

## Examples

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
