package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/lexsalg/goweb/config"
	"github.com/lexsalg/goweb/handlers"
	v1 "github.com/lexsalg/goweb/handlers/api/v1"
	u "github.com/lexsalg/goweb/utils"

	"github.com/gorilla/mux"
)

func main() {

	m := mux.NewRouter()

	m.HandleFunc("/", handlers.Index).Methods("GET")

	m.HandleFunc("/correo-else", handlers.App1).Methods("GET")
	m.HandleFunc(`/correo-else{url:[(\s|/)\w]+}`, handlers.App1).Methods("GET")
	m.HandleFunc("/tramite", handlers.App2).Methods("GET")
	m.HandleFunc(`/tramite{url:[(\s|/)\w]+}`, handlers.App2).Methods("GET")
	m.HandleFunc("/clima", handlers.App3).Methods("GET")
	m.HandleFunc(`/clima{url:[(\s|/)\w]+}`, handlers.App3).Methods("GET")

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

	static := http.FileServer(u.FileSystem{FS: http.Dir("static")})
	statics := http.StripPrefix(strings.TrimRight("/static/", "/"), static)
	m.PathPrefix("/static/").Handler(statics)

	apps := http.FileServer(u.FileSystem{FS: http.Dir("apps")})
	styles := http.StripPrefix(strings.TrimRight("/", "/"), apps)
	m.PathPrefix("/").Handler(styles)

	log.Println("running server, port:", config.ServerPort(), "...")
	server := &http.Server{
		Addr:    config.UrlServer(),
		Handler: m,
	}
	log.Fatal(server.ListenAndServe())
}
