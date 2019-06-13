package utils

import (
	"html/template"
	"net/http"

	"github.com/lexsalg/goweb/config"
)

var tpls = template.Must(template.New("t").ParseGlob(config.DirTemplate()))
var errTpl = template.Must(template.ParseFiles(config.DirTemplateError()))

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	if err := tpls.ExecuteTemplate(w, name, data); err != nil {
		RenderErrorTemplate(w, err)
	}
}

func RenderApp(w http.ResponseWriter, baseHref string) {
	dirApp := "apps/" + baseHref + "/index.html"
	ctx := make(map[string]interface{})
	ctx["BaseHref"] = baseHref

	w.Header().Set("Content-Type", "text/html")
	app, _ := template.Must(template.ParseFiles(dirApp)).ParseGlob("templates/help/*.html")
	if err := app.Execute(w, ctx); err != nil {
		RenderErrorTemplate(w, err)
	}
}

func RenderErrorTemplate(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = errTpl.Execute(w, err.Error())
}
