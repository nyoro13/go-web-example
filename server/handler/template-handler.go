package handler

import (
	"net/http"
	"strings"
)

type Lang int

const (
	_       = iota
	EN Lang = iota
	JA
)

func (lang Lang) String() string {
	switch lang {
	case EN:
		return "en"
	case JA:
		return "ja"
	default:
		return "Unknown"
	}
}

func (lang Lang) path() string {
	switch lang {
	case EN:
		return ""
	case JA:
		return "ja"
	default:
		return ""
	}
}

func convertPathToLang(path string) Lang {
	switch path {
	case EN.path():
		return EN
	case JA.path():
		return JA
	default:
		return EN
	}
}

type TemplateHandler struct {
	templateDir string
}

func MakeTemplateHandler(templateDir string) TemplateHandler {
	return TemplateHandler{templateDir: templateDir}
}

func (handler *TemplateHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	paths := strings.Split(request.URL.Path, "/")[1:]
	if paths[len(paths)-1] == "" {
		paths[len(paths)-1] = "index.html"
	}

	responseWriter.Write([]byte("hello"))
}
