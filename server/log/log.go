package log

import (
	"fmt"
	"log"
	"os"
)

var DebugLogger *log.Logger
var WarnLogger *log.Logger
var ErrorLogger *log.Logger

func SetStdLogger(prefix string) {
	DebugLogger = log.New(os.Stdout, prefix+" [DEBUG] ", log.LstdFlags)
	WarnLogger = log.New(os.Stdout, prefix+" [WARN] ", log.LstdFlags)
	ErrorLogger = log.New(os.Stdout, prefix+" [ERROR] ", log.LstdFlags)
}

func Debug(format string, a ...interface{}) {
	if DebugLogger != nil {
		DebugLogger.Println(fmt.Sprintf(format, a...))
	}
}

func Warn(format string, a ...interface{}) {
	if WarnLogger != nil {
		WarnLogger.Println(fmt.Sprintf(format, a...))
	}
}

func Error(format string, a ...interface{}) {
	if ErrorLogger != nil {
		ErrorLogger.Println(fmt.Sprintf(format, a...))
	}
}
