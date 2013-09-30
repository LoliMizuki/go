package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	TestChannel()
	// TestGoRoutine()
}

func TestChannel() {
	ch := make(chan string)
	go sender(ch)
	go recivie(ch)
	time.Sleep(1e9)
}

func sender(ch chan string) {
	ch <- "Yuyuko"
	ch <- "Pachi"
	ch <- "Ayaa"
	ch <- "Nekoko"
}

func recivie(ch chan string) {
	for {
		input := <-ch
		fmt.Println(input)
	}
}

func TestGoRoutine() {
	go PrintMessage("go_go_111", 2*1e9, 5) // 2 sec
	go PrintMessage("go_go_222", 4*1e9, 3) // 4 sec
	time.Sleep(10 * 1e9)                   // 10 sec
}

func PrintMessage(message string, interval time.Duration, times int) {
	for i := 0; i < times; i++ {
		fmt.Println("#", i, ": ", message)
		// time.Sleep(interval)
	}

	runtime.Goexit()
}
