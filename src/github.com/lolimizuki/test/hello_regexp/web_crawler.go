package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func web_crawler() {
	resp, err := http.Get("http://blog.golang.org/strings")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	htmlString := string(body)

	fmt.Println(htmlString)

	file, err := os.OpenFile("temp.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(htmlString)
}

func do_some_stuff() {
	sourceBytes := []byte(`<div>Maki chan kawaii!!!!</div><div>Yuyuko chan ... Pink is EROI</div>`)

	// re, err := regexp.Compile(`<[\S\s]+?>`)
	// re, err := regexp.Compile(`<div>([a-zA-Z]+)`)
	// re, err := regexp.Compile(`<div>[a-zA-Z\s[:graph:]]*</div>`)
	re, err := regexp.Compile(`<div>(?P<stmt>[a-zA-Z\s[:graph:]]*)</div>`)
	if err != nil {
		fmt.Println("compile error", err)
		return
	}

	sourceString := string(sourceBytes)

	// find
	findString := re.FindString(sourceString)
	fmt.Println("Find:", findString)

	// find all

	// find and to lower
	lowerString := re.ReplaceAllStringFunc(sourceString, func(src string) string {
		return strings.ToLower(src)
	})
	fmt.Println("ToLower:", lowerString)
}

func try_expnad() {
	src := []byte(`
        call hello alice
        hello bob
        call hello eve
    `)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
