package main

import (
	"github.com/lexsalg/goweb/rest/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	h "github.com/lexsalg/goweb/rest/handlers"
)

func main() {
	models.SetDefaultUser()

	m := mux.NewRouter()

	m.HandleFunc("/api/v1/users/", h.GetAll).Methods("GET")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", h.Get).Methods("GET")
	m.HandleFunc("/api/v1/users/", h.Create).Methods("POST")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", h.Update).Methods("PUT")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", h.Delete).Methods("DELETE")

	log.Println("running server port:8000...")
	log.Fatal(http.ListenAndServe("localhost:8000", m))
}
