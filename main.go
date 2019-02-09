package main

import (
	"go-web-example/log"
)

func main() {
	log.SetStdLogger("go-web-example")
	log.Debug("hello")
	log.Warn("world")
	log.Error("!!!")
}
