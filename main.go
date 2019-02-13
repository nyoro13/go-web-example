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

	/*
	appHandler := handler.NewAppHandler(server.GetTemplateDir(), server.GetStaticDir(), server.GetMessageDir())
	appHandler.AppendStaticPrefix("images")
	http.Handle("/", appHandler)
	*/

	staticHandler := handler.MakeStaticHandler(server.GetStaticDir())
	http.HandleFunc("/images/", staticHandler.Handle)
	http.HandleFunc("/scripts/", staticHandler.Handle)
	http.HandleFunc("/styles/", staticHandler.Handle)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
