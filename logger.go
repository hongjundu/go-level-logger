package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelNotice
	LogLevelWarning
	LogLevelError
)

const (
	LogLevelDebugTag   = "[DEBUG]"
	LogLevelInfoTag    = "[INFO]"
	LogLevelNoticeTag  = "[NOTICE]"
	LogLevelWarningTag = "[WRAN]"
	LogLevelErrorTag   = "[ERROR]"
)

const (
	LogTimeFormat = "2006-01-02 15:04:05:123.000"
)

var (
	once      sync.Once
	logLevel  int
	stdLogger *log.Logger
	errLogger *log.Logger
)

func InitLogger(level int) {
	logLevel = level
	once.Do(func() {
		stdLogger = log.New(os.Stdout, "", 0)
		errLogger = log.New(os.Stderr, "", 0)
	})
}

func InitLoggerWithOutput(stdOutput, errOutput io.Writer, level int) {
	logLevel = level
	once.Do(func() {
		stdLogger = log.New(stdOutput, "", 0)
		errLogger = log.New(errOutput, "", 0)
	})
}

func Debug(v ...interface{}) {
	if logLevel <= LogLevelDebug {
		logPrint(stdLogger, LogLevelDebugTag, v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel <= LogLevelDebug {
		logPrintf(stdLogger, LogLevelDebugTag, format, v...)
	}
}

func Debugln(v ...interface{}) {
	if logLevel <= LogLevelDebug {
		logPrintln(stdLogger, LogLevelDebugTag, v...)
	}
}

func Info(v ...interface{}) {
	if logLevel <= LogLevelInfo {
		logPrint(stdLogger, LogLevelInfoTag, v...)
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel <= LogLevelInfo {
		logPrintf(stdLogger, LogLevelInfoTag, format, v...)
	}
}

func Infoln(v ...interface{}) {
	if logLevel <= LogLevelInfo {
		logPrintln(stdLogger, LogLevelInfoTag, v...)
	}
}

func Notice(v ...interface{}) {
	if logLevel <= LogLevelNotice {
		logPrint(stdLogger, LogLevelNoticeTag, v...)
	}
}

func Noticef(format string, v ...interface{}) {
	if logLevel <= LogLevelNotice {
		logPrintf(stdLogger, LogLevelNoticeTag, format, v...)
	}
}

func Noticeln(v ...interface{}) {
	if logLevel <= LogLevelNotice {
		logPrintln(stdLogger, LogLevelNoticeTag, v...)
	}
}

func Warn(v ...interface{}) {
	if logLevel <= LogLevelWarning {
		logPrint(stdLogger, LogLevelWarningTag, v...)
		logPrint(errLogger, LogLevelWarningTag, v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if logLevel <= LogLevelWarning {
		logPrintf(stdLogger, LogLevelWarningTag, format, v...)
		logPrintf(errLogger, LogLevelWarningTag, format, v...)
	}
}

func Warnln(v ...interface{}) {
	if logLevel <= LogLevelWarning {
		logPrintln(stdLogger, LogLevelWarningTag, v...)
		logPrintln(errLogger, LogLevelWarningTag, v...)
	}
}

func Error(v ...interface{}) {
	if logLevel <= LogLevelError {
		logPrint(stdLogger, LogLevelErrorTag, v...)
		logPrint(errLogger, LogLevelErrorTag, v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel <= LogLevelError {
		logPrintf(stdLogger, LogLevelErrorTag, format, v...)
		logPrintf(errLogger, LogLevelErrorTag, format, v...)
	}
}

func Errorln(v ...interface{}) {
	if logLevel <= LogLevelError {
		logPrintln(stdLogger, LogLevelErrorTag, v...)
		logPrintln(errLogger, LogLevelErrorTag, v...)
	}
}

func Panic(v ...interface{}) {
	stdLogger.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	stdLogger.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	stdLogger.Panicln(v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func logPrint(logger *log.Logger, level string, v ...interface{}) {
	currtime := time.Now().Format(LogTimeFormat)
	newV := append([]interface{}{currtime, level}, v...)
	logger.Print(newV)
}

func logPrintf(logger *log.Logger, level string, format string, v ...interface{}) {
	currtime := time.Now().Format(LogTimeFormat)
	newV := append([]interface{}{currtime, level}, fmt.Sprintf(format, v...))
	logger.Println(newV)
}

func logPrintln(logger *log.Logger, level string, v ...interface{}) {
	currtime := time.Now().Format(LogTimeFormat)
	newV := append([]interface{}{currtime, level}, v...)
	logger.Println(newV)
}
