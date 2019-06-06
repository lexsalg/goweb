package main

import (
	"fmt"
	"github.com/lexsalg/goweb/projects/mux"
	"log"
	"net/http"
)

func hola(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hola desde una funcion anonima")
}

type User struct {
	name string
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hola usuario"+u.name)
}

func main() {
	user := &User{"alexis salgado"}
	m := mux.CreateMux()
	m.AddFunc("/hola", hola)
	m.AddHandle("/usuario", user)

	log.Fatal(http.ListenAndServe("localhost:3000", m))
}
