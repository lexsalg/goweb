package main

import (
	"fmt"
	"log"
	"net/http"
)

func Holax(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hola mundo")
}

func Holay(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hola mundo2")
}

func mainn() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Holax)
	http.HandleFunc("/", Holay) //DefaultServerMux

	server := &http.Server{
		Addr: "localhost:3000",
		//Handler: nil, //si es nil se utiliza DefaultServerMux
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
