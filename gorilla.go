package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	nombre := params["nombre"]
	id := params["id"]
	_, _ = fmt.Fprintln(w, "los parametros son "+nombre+" "+id)

}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/usuarios/{nombre}/{id:[0-9]+}", YourHandler).Methods("PUT", "DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}
