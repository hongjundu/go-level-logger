package logger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

// logLevel:
// 0->DEBUG
// 1->INFO
// 2->NOTICE
// 3->WARN
// 4->ERROR

// logFileName:
// typically application name is used

// logFileDir:
// the directory where log files dump to

// maxLogFileSize:
// in megabytes, 100 by default

// maxLogFileBackups:
// max file backups, 3 by default

// maxLogFileAge:
// max log file ages, in days. 30 by default

func Init(logLevel int, logFileName string, logFileDir string, maxLogFileSize, maxLogFileBackups, maxLogFileAge int) {
	if len(logFileDir) == 0 {
		InitLogger(logLevel)
		Infof("[logger] log path is not set")
		return
	}

	dirExists, _ := dirExists(logFileDir)

	if !dirExists {
		InitLogger(logLevel)
		Warnf("[logger] log path '%s' does NOT exist", logFileDir)
		return
	}

	maxSize := maxLogFileSize
	maxBackups := maxLogFileBackups
	maxAge := maxLogFileAge

	if maxSize <= 0 {
		maxSize = 100
	}

	if maxBackups <= 0 {
		maxBackups = 3
	}

	if maxAge <= 0 {
		maxAge = 30
	}

	stdOutputFile := &lumberjack.Logger{
		Filename:   filepath.Join(logFileDir, fmt.Sprintf("%s.log", logFileName)),
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, //days
	}

	errOutputFile := &lumberjack.Logger{
		Filename:   filepath.Join(logFileDir, fmt.Sprintf("%s-error.log", logFileName)),
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, //days
	}

	stdOutput := io.MultiWriter(stdOutputFile, os.Stdout)
	errOutput := io.MultiWriter(errOutputFile, os.Stderr)

	InitLoggerWithOutput(stdOutput, errOutput, logLevel)
}

func dirExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else {
		if os.IsNotExist(err) {
			return false, err
		} else {
			return true, err
		}
	}
}
