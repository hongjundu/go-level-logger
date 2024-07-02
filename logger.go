package logger

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	LogLevelWarningTag = "[WARN]"
	LogLevelErrorTag   = "[ERROR]"
)

const (
	RotateStrategyPerHour = iota + 1
	RotateStrategyPerDay
)

const (
	LogTimeFormat = "2006-01-02 15:04:05.000"
)

var (
	once              sync.Once
	logLevel          int
	stdLogger         *log.Logger
	errLogger         *log.Logger
	stdOutputFile     *lumberjack.Logger
	errOutputFile     *lumberjack.Logger
	rotateLog         bool
	rotateLogStrategy int
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

		if rotateLog {
			c := cron.New(cron.WithSeconds())
			spec := "0 0 0 * * ?"
			if rotateLogStrategy == RotateStrategyPerHour {
				spec = "0 0 */1 * * ?"
			}
			//rolling strategy: rotate at 00:00:00 every day
			c.AddFunc(spec, func() {
				stdOutputFile.Rotate()
				errOutputFile.Rotate()
			})
			c.Start()
		}
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
	logPrintln(errLogger, LogLevelErrorTag, v...)
	log.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	logPrintf(errLogger, LogLevelErrorTag, format, v...)
	log.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func logPrint(logger *log.Logger, level string, v ...interface{}) {
	currtime := time.Now().Format(LogTimeFormat)
	fileinfo := fileline()
	newV := append([]interface{}{currtime, fileinfo, level}, v...)
	logger.Print(newV)
}

func logPrintf(logger *log.Logger, level string, format string, v ...interface{}) {
	currtime := time.Now().Format(LogTimeFormat)
	fileinfo := fileline()
	newV := append([]interface{}{currtime, fileinfo, level}, fmt.Sprintf(format, v...))
	logger.Println(newV)
}

func logPrintln(logger *log.Logger, level string, v ...interface{}) {
	currtime := time.Now().Format(LogTimeFormat)
	fileinfo := fileline()
	newV := append([]interface{}{currtime, fileinfo, level}, v...)
	logger.Println(newV)
}

func fileline() string {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		file = filepath.Base(file)
	}
	if len(file) == 0 {
		file = "???"
	}
	if line < 0 {
		line = 0
	}

	return fmt.Sprintf("%s:%d", file, line)
}
