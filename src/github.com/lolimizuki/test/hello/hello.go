package main

import (
	"fmt"
)

func main() {
	hello := "Hello World"

	for _, c := range hello {
		fmt.Println(string(c))
	}
}
