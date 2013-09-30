package main

import (
	"fmt"
)

type ThisMethods interface {
	Add(x int)
	Get() int
}

type Parent struct {
	ParentFiled int
}

func (parent *Parent) ParentMethod(v int) int {
	parent.ParentFiled += v
	return parent.ParentFiled
}

func (Parent) Stuff() {
	fmt.Println("Parent Stuff ... ")
}

func (self Parent) CallStuff() {
	self.Stuff()
	self.Stuff()
}

type Child struct {
	Parent // inheritance
	// ThisMethods // interface
	// ChildFiled  int
}

func (child *Child) Add(x int) {
	// child.ChildFiled = child.ParentFiled + x
}

func (child *Child) Get() int {
	// return child.ChildFiled
	return 999
}

// Override
func (Child) Stuff() {
	fmt.Println("Child Stuff")
}

type Base struct{}

func (Base) Magic() { fmt.Print("base magic") }
func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() { fmt.Println("voodoo magic") }

func main() {
	// c := new(Child)
	// c.ParentMethod(12) // 可以直接呼叫 Parent 的方法

	// fmt.Println(c.ParentFiled)

	// c.Add(99)
	// fmt.Println(c.Get())

	// c.CallStuff()

	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
