package main

import (
	"fmt"
)

type Displayer interface {
	Show()
}

type MyStruct struct {
}

func (m *MyStruct) Show() {
	fmt.Println("I am MyStruct >...<")
}

func main() {
	m := MyStruct{}

	// var vi interface{} = m // can not test ... :D
	var vi interface{} = &m

	if mi, ok := vi.(Displayer); ok {
		mi.Show()
	}
}
