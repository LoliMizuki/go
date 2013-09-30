package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// use_os_arg()
	// use_flag()
	// create_test_file("1", "test1111.txt")
	// create_test_file("2", "test2222.txt")
}

func create_test_file(s string, filename string) {
	totalLine := 10
	outputFile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer outputFile.Close()

	for i := 0; i < totalLine; i++ {
		for j := 0; j < i+1; j++ {
			outputFile.WriteString(s)
		}
		outputFile.WriteString("\n")
	}
}

func use_os_arg() {
	for i, a := range os.Args {
		fmt.Println(i, ": ", a)
	}
}

func use_flag() {
	flag.PrintDefaults()
	flag.Parse()

	for i, arg := range flag.Args() {
		fmt.Println(i, "::", arg)
	}
}