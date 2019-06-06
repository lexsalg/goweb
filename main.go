package main

import (
	"html/template"
	"log"
	"net/http"
)

//var tpls = template.Must(template.New("T").ParseFiles(
//	"templates/index.html",
//	"templates/footer.html",
//	"templates/header.html",
//	"templates/registro.html",
//))

var tpls = template.Must(template.New("T").ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpls.ExecuteTemplate(w, "registro", nil)
		if err != nil {
			panic(err)
		}
	})
	log.Println("el servidor escucha:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
