package common

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

type Comment struct {
	Title       string
	Content     string
	ContentHtml template.HTML
}

var tpls = template.Must(template.New("T").ParseGlob("templates/**/*.html"))

var errTpl = template.Must(template.ParseFiles("templates/error.html"))

func handlerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = errTpl.Execute(w, err.Error())
}

func renderTpl(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := tpls.ExecuteTemplate(w, name, data)
	if err != nil {
		handlerError(w, r, err)
	}
}

func CommentController(w http.ResponseWriter, r *http.Request) {

	contentHtml := template.HTML("<h1>contenido</h1>")
	c := Comment{"titulo",
		"<h1>contenido</h1>",
		contentHtml,
	}
	renderTpl(w, r, "comments/comment", c)
}
