package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main6() {
	URL := createURL() //url partiendo de una uri
	fmt.Println(URL)
	//http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(r.Header)
	//	accessToken := r.Header.Get("access_token")
	//	if len(accessToken) != 0 {
	//		fmt.Println(accessToken)
	//	}
	//
	//	r.Header.Set("nombreheader", "valorheader")
	//	fmt.Println(r.Header)
	//	_, _ = fmt.Fprintln(w, "hola mundo")
	//})
	//
	//log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func main5() {
	URL := createURL() //url partiendo de una uri
	fmt.Println(URL)
}
func createURL() string {
	u, err := url.Parse("/params")
	if err != nil {
		panic(err)
	}
	u.Host = "localhost:3000"
	u.Scheme = "http"
	query := u.Query()
	query.Add("nombre", "valor")
	u.RawQuery = query.Encode()
	return u.String()
}

func main4() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		accessToken := r.Header.Get("access_token")
		if len(accessToken) != 0 {
			fmt.Println(accessToken)
		}

		r.Header.Set("nombreheader", "valorheader")
		fmt.Println(r.Header)
		_, _ = fmt.Fprintln(w, "hola mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
func main3() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.URL)
		values := r.URL.Query()
		values.Del("otro	")
		values.Add("name", "alexis")
		values.Add("name1", "alexis1")
		values.Add("name2", "alexis2")
		r.URL.RawQuery = values.Encode()
		fmt.Println(r.URL)
		_, _ = fmt.Fprintln(w, "hola mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func main2() {
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.RawQuery) // solo muestra la uri, sin dominio,query es ?
		fmt.Println(r.URL.Query()) // retorna mapa
		name := r.URL.Query().Get("name")
		if len(name) != 0 {
			fmt.Println(name)
		}

		_, _ = fmt.Fprintln(w, "hola mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
