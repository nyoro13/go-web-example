package main

import (
	"go-web-example/server/handler"
	"go-web-example/server"
	"go-web-example/server/log"
	"net/http"
)

func main() {
	log.SetStdLogger("GoWebExample")

	option := server.ParseOption()
	server.InitConfig(option.Mode)

	log.Debugfln("Mode: %s", server.GetMode())

	appHandler := handler.NewAppHandler("./client/templates", "./client", "./client/messages")
	appHandler.AppendStaticPrefix("images")
	http.Handle("/", appHandler)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
