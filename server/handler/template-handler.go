package handler

import (
	"net/http"
)

type TemplateHandler struct {
	templateDir string
}

func (handler *TemplateHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
}
