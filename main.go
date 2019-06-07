package main

import (
	"github.com/lexsalg/goweb/common"
	"log"
	"net/http"
)

func main() {

	staticFiles := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.HandleFunc("/comment/", common.CommentController)

	mux.Handle("/assets/", http.StripPrefix("/assets/", staticFiles))

	log.Println("running server port:3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
