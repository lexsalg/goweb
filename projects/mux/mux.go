package mux

import (
	"net/http"
)

type customHandler func(http.ResponseWriter, *http.Request)

type AlexisMux struct {
	rutasAlexis map[string]customHandler //handlers
}

func (m *AlexisMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fn, ok := m.rutasAlexis[r.URL.Path]; ok {
		fn(w, r)
	} else {
		http.NotFound(w, r)
	}

}

func (m *AlexisMux) AddFunc(ruta string, fn customHandler) {
	m.rutasAlexis[ruta] = fn
}

func (m *AlexisMux) AddHandle(ruta string, handle http.Handler) {
	m.rutasAlexis[ruta] = handle.ServeHTTP
}

func CreateMux() *AlexisMux {
	newMap := make(map[string]customHandler)
	return &AlexisMux{newMap}

}
