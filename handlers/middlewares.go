package handlers

import (
	"fmt"
	"net/http"

	"github.com/lexsalg/goweb/utils"
)

type customHandler func(w http.ResponseWriter, r *http.Request)

func Authentication(function customHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsAuth(r) {
			http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			return
		}
		function(w, r)
	})
}

func MiddlewareTwo(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("deberia borrar cache")
		// w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		fmt.Println("yanoda")
		handler.ServeHTTP(w, r)
	})
}
