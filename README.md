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
        "gopkg.in/natefinch/lumberjack.v2"
        "io"
    )


    type compWriter struct {
        writers []io.Writer
    }

    func (cw *compWriter) Write(p []byte) (n int, err error) {
        for _, w := range cw.writers {
            n, err = w.Write(p)
        }
        return
    }

    func newCompWriter(writers ...io.Writer) io.Writer {
        return &compWriter{writers}
    }

    func main() {
        stdOutputFile := &lumberjack.Logger{
            Filename:   filepath.Join(utils.GetLogsDir(), "invoicemgr.log"),
            MaxSize:    100, // megabytes
            MaxBackups: 3,
            MaxAge:     365, //days
        }

        errOutputFile := &lumberjack.Logger{
            Filename:   filepath.Join(utils.GetLogsDir(), "invoicemgr-error.log"),
            MaxSize:    100, // megabytes
            MaxBackups: 3,
            MaxAge:     365, //days
        }

        stdOutput := newCompWriter(stdOutputFile, os.Stdout)
        errOutput := newCompWriter(errOutputFile, os.Stderr)

        logger.InitLoggerWithOutput(stdOutput, errOutput, 0)

        logger.Debugf("debug message")
        logger.Infof("info message")
        logger.Noticef("notice message")
        logger.Warnf("warn message")
        logger.Errorf("error message")
    }
