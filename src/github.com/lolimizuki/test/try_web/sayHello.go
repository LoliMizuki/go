package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintf(w, "hello nekoko: ")

	fmt.Println("URL.Path: ", r.URL.Path)
	fmt.Println("r.URL.Scheme: ", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println("key:", k, "value:", v)
	}
}
