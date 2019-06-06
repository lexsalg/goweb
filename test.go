package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "hola mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
