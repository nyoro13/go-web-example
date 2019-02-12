package log

import (
	"fmt"
	"log"
	"os"
)

// DebugLogger is logger for debug
var DebugLogger *log.Logger

// WarnLogger is logger for warn
var WarnLogger *log.Logger

// ErrorLogger is logger for error
var ErrorLogger *log.Logger

// SetStdLogger set loggers as std out, err
func SetStdLogger(prefix string) {
	DebugLogger = log.New(os.Stdout, prefix+" [DEBUG] ", log.LstdFlags)
	WarnLogger = log.New(os.Stdout, prefix+" [WARN] ", log.LstdFlags)
	ErrorLogger = log.New(os.Stdout, prefix+" [ERROR] ", log.LstdFlags)
}

func Debugln(v ...interface{}) {
	if DebugLogger != nil {
		DebugLogger.Println(v...)
	}
}

func Debugfln(format string, v ...interface{}) {
	Debugln(fmt.Sprintf(format, v...))
}

func Warnln(v ...interface{}) {
	if WarnLogger != nil {
		WarnLogger.Println(v...)
	}
}

func Warnfln(format string, v ...interface{}) {
	Warnln(fmt.Sprintf(format, v...))
}

func Errorln(v ...interface{}) {
	if ErrorLogger != nil {
		ErrorLogger.Println(v...)
	}
}

func Errorfln(format string, v ...interface{}) {
	Errorln(fmt.Sprintf(format, v...))
}

func Fatalln(v ...interface{}) {
	Errorln(v...)
	os.Exit(1)
}

func Fatalfln(format string, v ...interface{}) {
	Fatalln(fmt.Sprintf(format, v...))
}
