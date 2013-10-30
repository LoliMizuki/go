package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func loginWithCookie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie, err := r.Cookie("username")

		if err == http.ErrNoCookie {
			t, _ := template.ParseFiles("login_cookie.gtpl")
			fmt.Println(err)
			t.Execute(w, nil)
		} else {
			fmt.Fprintf(w, "Good, you have login my page, your cookie is %v", cookie)
		}
	} else {
		if r.Form.Get("password") != r.Form.Get("username") && r.Form.Get("username") != "mizuki" {
			fmt.Fprintf(w, "login fail, username must be equal to password (??")
			return
		}

		now := time.Now()
		duration := time.Duration(time.Hour * 3)
		expiration := now.Add(duration)
		cookie := http.Cookie{Name: "username", Value: "mizuki", Expires: expiration}

		http.SetCookie(w, &cookie)

		fmt.Fprintf(w, "YES, you login this page, please try to refresh :D\n")
	}
}
