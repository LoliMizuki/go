package main

import (
	"fmt"
)

type MyInterface interface {
	Foo()
	Koo()
}

type FakeSuper struct {
	Field int
}

type MyStruct struct {
	MyInterface
	*FakeSuper
}

func (m *MyStruct) Foo() {
	fmt.Println("Foo")
}

func (m *MyStruct) Koo() {
	fmt.Println(m.Field)
}

func main() {
	f := &FakeSuper{234}
	m := &MyStruct{}
	m.FakeSuper = f

	m.Foo()
	m.Koo()
}
