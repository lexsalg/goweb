package main

import (
	"fmt"
	"log"
	"net/http"
)

//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}

type User struct {
	name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hola:"+this.name)
}

func main2() {
	alexis := &User{name: "Alexis"}
	mux := http.NewServeMux()
	mux.Handle("/alexis", alexis)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
