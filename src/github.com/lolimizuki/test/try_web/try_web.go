package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/user_input", userInput)
	http.HandleFunc("/upload_file", uploadFile)
	http.HandleFunc("/login_cookie", loginWithCookie)
	// http.HandleFunc("/login_session", loginWithSession)
	http.HandleFunc("/private", nekoPrivateGarden)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func nekoPrivateGarden(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	fmt.Fprintf(w, "this is mimi's private garden")
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("Method", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Fprintf(w, "User Name: %s\n", r.Form["username"])
		fmt.Fprintf(w, "Password: %s\n", r.Form["password"])

		// 輸入驗證 ...
		ageString := r.Form["age"][0]
		fmt.Println(ageString)
		if matchOk, _ := regexp.MatchString("^[0-9]+$", ageString); matchOk {
			fmt.Fprintf(w, "Age: %s", ageString)
		} else {
			fmt.Fprintf(w, "You suck not give me number ... !!!!")
		}

		fmt.Fprintf(w, "\n")

		// drop menu
		selectTerms := []string{"apple", "pear", "banane"}
		userSelectedTerm := r.Form.Get("fruit")
		findInMyTerms := false
		for _, term := range selectTerms {
			if userSelectedTerm == term {
				findInMyTerms = true
				fmt.Fprintf(w, "You pick up %s", term)
				break
			}
		}
		if findInMyTerms == false {
			fmt.Fprintf(w, "What are you send to me??? %s", userSelectedTerm)
		}

		// radio
		fmt.Fprintf(w, "\n")
		radioSelection := r.Form.Get("girls")
		radioTerms := []string{"1", "2", "3"}
		foundInRadioTerm := false
		for _, radioTerm := range radioTerms {
			if radioTerm == radioSelection {
				foundInRadioTerm = true
				break
			}
		}
		if foundInRadioTerm == true {
			fmt.Fprintf(w, "FOUND!!!! you are select is %s", radioSelection)
		} else {
			fmt.Fprintf(w, "OOPS NOT FOUND!!!! you are select is %s", radioSelection)
		}

		// checkbox multi selection
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "%s", r.Form["sleeplist"])
	}
}
