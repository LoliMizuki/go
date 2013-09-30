package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	Go()
	Down()
}

type MyTagType struct {
	MyInterface
	ii int     "My int"
	ff float64 "My float"
	ss string  "My string"
}

func main() {
	var mm MyTagType
	printAllTags(mm)

	p_mm := &MyTagType{}
	p_mm.ss = "123"

	mm.DoSomething()
	fmt.Println(mm.ii)
	fmt.Println(p_mm.ss)
}

func printAllTags(mm MyTagType) {
	for i := 0; i < 4; i++ {
		mmType := reflect.TypeOf(mm)
		tag_of_field_i := mmType.Field(i)
		fmt.Printf("%v\n", tag_of_field_i.Tag)
	}
}

// MyTagType's method
func (m *MyTagType) DoSomething() {
	m.ii = 100
	m.ff = 3.14
	m.ss = "MyTagType"
}

func (m *MyTagType) SetIi(newIi int) {
	m.ii = newIi
}
