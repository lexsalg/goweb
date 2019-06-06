package main

import (
	"fmt"
	"log"
	"net/http"
)

type customHandler func(http.ResponseWriter, *http.Request)

type MuxAlexis struct {
	rutasAlexis map[string]customHandler //handlers
}

func (m *MuxAlexis) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := m.rutasAlexis[r.URL.Path]
	fn(w, r)
}

func (m *MuxAlexis) AddMux(ruta string, fn customHandler) {
	m.rutasAlexis[ruta] = fn
}
func main() {
	newMap := make(map[string]customHandler)
	mux := &MuxAlexis{rutasAlexis: newMap}
	mux.AddMux("/hola", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hola desde una funcion anonima")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}

//func main() {
//
//	redirect := http.RedirectHandler("http://www.dev-bug.com", http.StatusMovedPermanently)
//	notFound := http.NotFoundHandler()
//
//	//http.Handle("/redirect", redirect)
//	//http.Handle("/not", notFound)
//	//
//	//server := &http.Server{
//	//	Addr:    "localhost:3000",
//	//	Handler: nil,
//	//}
//	mux := http.NewServeMux()
//	mux.Handle("/redirect", redirect)
//	mux.Handle("/not", notFound)
//
//	server := &http.Server{
//		Addr:    "localhost:3000",
//		Handler: mux,
//	}
//
//	log.Fatal(server.ListenAndServe())
//}
