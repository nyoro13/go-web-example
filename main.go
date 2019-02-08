package main

import (
	"go-web-example/weblog"
)

func main() {
	logger := weblog.NewDefaultLogger("go-web-example")
	logger.Debug("Hello")
	logger.Warn("World")
	logger.Error("!!!")
}
