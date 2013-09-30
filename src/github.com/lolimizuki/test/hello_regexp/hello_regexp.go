package main

import (
	"fmt"
	"regexp"
)

func main() {
	soucerString := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+.[0-9]+"

	if ok, _ := regexp.Match(pattern, []byte(soucerString)); ok { // second return is error
		fmt.Println("Found")
	}

	re, _ := regexp.Compile(pattern)
	afterReplace := re.ReplaceAllString(soucerString, "##.#")
	fmt.Println(afterReplace)
}
