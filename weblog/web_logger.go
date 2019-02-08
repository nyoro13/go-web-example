package weblog

import (
	"fmt"
	"log"
	"os"
)

type WebLogger struct {
	debugLogger *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func NewDefaultLogger(prefix string) *WebLogger {
	return &WebLogger{
		debugLogger: log.New(os.Stdout, prefix+" [DEBUG] ", log.LstdFlags),
		warnLogger:  log.New(os.Stdout, prefix+" [WARN] ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, prefix+" [ERROR] ", log.LstdFlags),
	}
}

func (logger *WebLogger) Debug(format string, a ...interface{}) {
	logger.debugLogger.Println(fmt.Sprintf(format, a...))
}

func (logger *WebLogger) Warn(format string, a ...interface{}) {
	logger.warnLogger.Println(fmt.Sprintf(format, a...))
}

func (logger *WebLogger) Error(format string, a ...interface{}) {
	logger.errorLogger.Println(fmt.Sprintf(format, a...))
}
