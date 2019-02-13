package handler

import (
	"go-web-example/server/file"
	"net/http"
)

type StaticHandler struct {
	fileDir    string
	fileReader file.FileReader
}

func MakeStaticHandler(fileDir string) StaticHandler {
	return StaticHandler{fileDir: fileDir, fileReader: file.MakeFileReader()}
}

func (handler *StaticHandler) Handle(response http.ResponseWriter, request *http.Request) {
	data, err := handler.fileReader.Read(handler.fileDir + request.URL.Path)
	if err != nil {
		response.WriteHeader(404)
		response.Write([]byte("NotFound!"))
		return
	}
	response.Write(data)
}
