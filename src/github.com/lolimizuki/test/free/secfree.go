package main

import (
	"fmt"
	"time"
)

func FuncCanExportFromSecFree() string {
	return "I am in second free ><"
}

func TimeAdd() {
	now := time.Now()
	afterThreeSecond := now.Add(time.Duration(time.Second * 3))
	fmt.Println(now, "-->", afterThreeSecond)
}
