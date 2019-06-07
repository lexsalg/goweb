package main

import (
	"html/template"
	"log"
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

func handlerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = errTpl.Execute(w, err.Error())
}

func renderTpl(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := tpls.ExecuteTemplate(w, name, data)
	if err != nil {
		handlerError(w, err)
	}
}

func comment(w http.ResponseWriter, r *http.Request) {

	contentHtml := template.HTML("<h1>contenido</h1>")
	c := Comment{"titulo",
		"<h1>contenido</h1>",
		contentHtml,
	}
	renderTpl(w, "comments/comment", c)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTpl(w, "index", nil)
	})

	user := User{"alexis"}
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		renderTpl(w, "user", user)
	})

	http.HandleFunc("/comment", comment)
	//c := Comment{"titulo", "contenido"}
	//http.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
	//	renderTpl(w, "comments/comment", c)
	//})

	log.Println("el servidor escucha:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
