package utils

import (
	"github.com/lexsalg/goweb/rest/config"
	"html/template"
	"net/http"
)

var tpls = template.Must(template.New("t").ParseGlob(config.DirTemplate()))
var errTpl = template.Must(template.ParseFiles(config.DirTemplateError()))

func RenderErrorTemplate(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = errTpl.Execute(w, err.Error())
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := tpls.ExecuteTemplate(w, name, data)
	if err != nil {
		RenderErrorTemplate(w, err)
	}
}
