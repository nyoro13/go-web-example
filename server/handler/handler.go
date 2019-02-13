package handler

import (
	"encoding/json"
	"go-web-example/server/file"
	"go-web-example/server/util/slice"
	"net/http"
	"strings"
)

type AppHandler struct {
	templateDir    string
	staticDir      string
	messageDir     string
	staticPrefixes []string
	fileReader     file.FileReader
}

func NewAppHandler(templateDir string, staticDir string, messageDir string) *AppHandler {
	handler := new(AppHandler)
	handler.templateDir = templateDir
	handler.staticDir = staticDir
	handler.messageDir = messageDir
	handler.staticPrefixes = make([]string, 3)
	handler.fileReader = file.MakeFileReader()

	return handler
}

func (app *AppHandler) AppendStaticPrefix(prefix string) {
	app.staticPrefixes = append(app.staticPrefixes, prefix)
}

func (app AppHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	paths := strings.Split(request.URL.Path, "/")[1:]

	if paths[len(paths)-1] == "" {
		paths[len(paths)-1] = "index.html"
	}

	if app.isStaticPath(paths[0]) {
		template, err := app.fileReader.Read(app.staticDir + request.URL.Path)
		if err != nil {
			response.WriteHeader(404)
			response.Write([]byte("Not Found!!!"))
			return
		}
		response.Write(template)
		return
	}

	filename := paths[len(paths)-1]
	if strings.HasSuffix(filename, ".html") {
		paths[len(paths)-1] = filename[:len(filename)-5]
	}

	messageBin, err := app.fileReader.Read(app.messageDir + "/" + strings.Join(paths, "/") + "-en.json")
	if err != nil {
		response.WriteHeader(404)
		response.Write([]byte("Not Found!!!"))
		return
	}

	var message map[string]interface{}
	if err = json.Unmarshal(messageBin, &message); err != nil {
		response.WriteHeader(404)
		response.Write([]byte("Not Found!!!"))
		return
	}

	response.Write([]byte("Hoge"))
}

func (app *AppHandler) isStaticPath(firstPath string) bool {
	return slice.Contains(app.staticPrefixes, firstPath)
}
