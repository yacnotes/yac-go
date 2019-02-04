package log

import (
	"log"
	"os"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
	LevelInitFail
)

var Level int
var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}

func Debug(a ...interface{}) {
	if Level <= LevelDebug {
		logger.SetPrefix("[DEBUG] ")
		logger.Println(a...)
	}
}

func Info(a ...interface{}) {
	if Level <= LevelInfo {
		logger.SetPrefix("[INFO]  ")
		logger.Println(a...)
	}
}

func Warn(a ...interface{}) {
	if Level <= LevelWarn {
		logger.SetPrefix("[WARN]  ")
		logger.Println(a...)
	}
}

func Error(a ...interface{}) {
	if Level <= LevelError {
		logger.SetPrefix("[ERROR] ")
		logger.Println(a...)
	}
}

func Panic(a ...interface{}) {
	if Level <= LevelPanic {
		logger.SetPrefix("[PANIC] ")
		logger.Panicln(a...)
	}
}

func InitFail(a ...interface{}) {
	if Level <= LevelInitFail {
		logger.SetPrefix("[INIT FAIL] ")
		logger.Fatalln(a...)
	}
}

func GetLogger() *log.Logger {
	return logger
}
