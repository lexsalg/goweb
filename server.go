package goweb

import (
	"github.com/lexsalg/goweb/config"
	"github.com/lexsalg/goweb/handlers"
	"github.com/lexsalg/goweb/handlers/api/v1"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	m := mux.NewRouter()

	m.HandleFunc("/", handlers.Index)

	m.HandleFunc("/api/v1/users/", v1.GetAll).Methods("GET")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.Get).Methods("GET")
	m.HandleFunc("/api/v1/users/", v1.Create).Methods("POST")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.Update).Methods("PUT")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.Delete).Methods("DELETE")

	log.Println("running server, port:", config.ServerPort(), "...")

	server := &http.Server{
		Addr:    config.UrlServer(),
		Handler: m,
	}
	log.Fatal(server.ListenAndServe())
}
