package main

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

	m.HandleFunc("/", handlers.Index).Methods("GET")
	m.HandleFunc("/users/new", handlers.NewUser).Methods("GET", "POST")
	//m.HandleFunc("/users/edit", handlers.UpdateUser).Methods("GET")
	editHandler := handlers.Authentication(handlers.UpdateUser)
	editHandler = handlers.MiddlewareTwo(editHandler)

	m.Handle("/users/edit", editHandler).Methods("GET")

	m.HandleFunc("/users/login", handlers.Login).Methods("GET", "POST")
	m.HandleFunc("/users/logout", handlers.Logout).Methods("GET")

	m.HandleFunc("/api/v1/users/", v1.GetAll).Methods("GET")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.Get).Methods("GET")
	m.HandleFunc("/api/v1/users/", v1.Create).Methods("POST")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.Update).Methods("PUT")
	m.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.Delete).Methods("DELETE")

	assets := http.FileServer(http.Dir(config.DirAssets()))
	statics := http.StripPrefix(config.PrefixAssets(), assets)
	m.PathPrefix(config.PrefixAssets()).Handler(statics)

	log.Println("running server, port:", config.ServerPort(), "...")
	server := &http.Server{
		Addr:    config.UrlServer(),
		Handler: m,
	}
	log.Fatal(server.ListenAndServe())
}
