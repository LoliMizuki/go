package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func userInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("user_input.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		usUserInputText := r.Form.Get("inputtext")
		fmt.Println("usUserInputText:", usUserInputText)
		sUserInputText := template.HTMLEscapeString(usUserInputText)

		fmt.Fprintf(w, "You input: \n")
		fmt.Fprintf(w, "%s", sUserInputText)
	}
}
