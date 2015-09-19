package main

import (
	"fmt"
)

func MainInThird() {
	newNil()
	newSlice()
}

type YouStruct struct {
	sss []byte
	kkk int
}

func newNil() {
	// it's become panic error at runtime
	// var ys *YouStruct
	// fmt.Println(ys.kkk)
}

func newSlice() {
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = i
	}

	slice := array[2:4:7]
	fmt.Println(slice)
	fmt.Println(cap(slice))
}
