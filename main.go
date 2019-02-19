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

	staticHandler := handler.MakeStaticHandler(server.GetStaticDir())
	http.HandleFunc("/images/", staticHandler.Handle)
	http.HandleFunc("/scripts/", staticHandler.Handle)
	http.HandleFunc("/styles/", staticHandler.Handle)
	http.HandleFunc("/favicon.ico", staticHandler.Handle)

	topHandler := handler.MakeTemplateHandler(server.GetTemplateDir(), server.GetMessageDir())
	hogeHandler := handler.MakeTemplateHandler(server.GetTemplateDir(), server.GetMessageDir())
	http.HandleFunc("/ja/hoge/", hogeHandler.Handle)
	http.HandleFunc("/hoge/", hogeHandler.Handle)
	http.HandleFunc("/ja/", topHandler.Handle)
	http.HandleFunc("/", topHandler.Handle)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
