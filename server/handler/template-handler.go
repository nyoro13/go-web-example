package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
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

func convertPathToLang(langPath string) Lang {
	switch langPath {
	case EN.path():
		return EN
	case JA.path():
		return JA
	default:
		return EN
	}
}

func parsePath(path string) (basePass string, lang Lang) {
	paths := strings.Split(path, "/")[1:]
	if paths[len(paths)-1] == "" {
		paths[len(paths)-1] = "index.html"
	}

	lang = convertPathToLang(paths[0])
	if paths[0] == lang.path() {
		paths = paths[1:]
	}

	basePass = "/" + strings.Join(paths, "/")
	return
}

type TemplateHandler struct {
	templateDir string
	messageDir  string
	template    *template.Template
	messages    map[Lang]interface{}
}

func MakeTemplateHandler(templateDir string, messageDir string) TemplateHandler {
	return TemplateHandler{
		templateDir: templateDir,
		messageDir:  messageDir,
	}
}

func (handler *TemplateHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	basePath, lang := parsePath(request.URL.Path)
	if handler.template == nil {
		if handler.template = handler.getTemplate(basePath); handler.template == nil {
			responseWriter.WriteHeader(404)
			responseWriter.Write([]byte("Not Found"))
			return
		}
	}

	message, ok := handler.messages[lang]
	if !ok {
		message = handler.getMessage(basePath, lang)
		handler.messages[lang] = message
	}

	handler.template.ExecuteTemplate(responseWriter, basePath, message)
}

func (handler *TemplateHandler) getTemplate(basePath string) *template.Template {
	return template.Must(template.ParseFiles(handler.templateDir + basePath))
}

func (handler *TemplateHandler) getMessage(basePath string, lang Lang) interface{} {
	var messagePath string
	if strings.HasSuffix(basePath, ".html") {
		messagePath = basePath[:len(basePath)-5]
	}
	messagePath += "-" + lang.String() + ".json"
	messageJSON, err := ioutil.ReadFile(messagePath)
	if err != nil {
		handler.messages[lang] = nil
	}

	// TODO 型指定できないやんか
	var hoge interface{}
	err = json.Unmarshal(messageJSON, &hoge)
	return hoge
}
