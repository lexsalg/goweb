package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func createURL2() string {
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

func main() {

	URL := createURL2()
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println("el header es:", response.Header)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("el body es:", string(body))
	fmt.Println("el status es:", response.Status)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
