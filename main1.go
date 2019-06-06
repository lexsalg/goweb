package main

import (
	"fmt"
	"log"
	"net/http"
)

func main1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("el metodo es:", r.Method)
		w.Header().Add("alexis", "prueba de heder")
		//_, _ = fmt.Fprintln(w, "hola mundo")
		//http.Redirect(w, r, "/dos", http.StatusMovedPermanently)
		http.Redirect(w, r, "https://www.dev-bug.com", http.StatusMovedPermanently)
	})

	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "hola mundo", "dos")
	})

	http.HandleFunc("/tres", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ESTE ES UN ERROR", http.StatusConflict)
	})

	http.HandleFunc("/method", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			_, _ = fmt.Fprintln(w, "GET")
		case "POST":
			_, _ = fmt.Fprintln(w, "POST")
		case "PUT":
			_, _ = fmt.Fprintln(w, "PUT")
		case "DELETE":
			_, _ = fmt.Fprintln(w, "DELETE")
		default:
			http.Error(w, "metodo no valido", http.StatusBadRequest)
		}

		//_, _ = fmt.Fprintln(w, "hola mundo")

	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
