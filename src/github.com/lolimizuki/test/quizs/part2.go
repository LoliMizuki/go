package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func Q24_Interface() {
	fmt.Println(g(123))
	fmt.Println(g("123"))
}

func g(anything interface{}) int {
	if v, ok := anything.(int); ok {
		return v
	}

	return -1
}

func Q26_Max() {
	ints := []int{1, 2, 3, 4, 5}
	findMaxInInterfaceSlice(interfacesFromInts(ints)) // 無法傳 []int 給 []interface{}??? use suck solution

	float64s := []float64{7.123, 1.1, 3.2, 4.6, 2.2}
	findMaxInInterfaceSlice(interfacesSliceFromFloat64s(float64s))
}

func interfacesFromInts(ints []int) []interface{} {
	interfaces := make([]interface{}, len(ints))
	for i, v := range ints {
		interfaces[i] = v
	}

	return interfaces
}

func interfacesSliceFromFloat64s(float64s []float64) []interface{} {
	interfaces := make([]interface{}, len(float64s))
	for i, v := range float64s {
		interfaces[i] = v
	}

	return interfaces
}

func findMaxInInterfaceSlice(anyTypes []interface{}) {
	if len(anyTypes) == 0 {
		fmt.Println("Empty!!!")
		return
	}

	isGreatThan := func(testValue interface{}) func(ix, iy interface{}) bool {
		switch testValue.(type) {
		case int:
			return func(ix, iy interface{}) bool {
				if _, ok := iy.(int); ok {
					return ix.(int) > iy.(int)
				}

				return false
			}

		case float64:
			return func(ix, iy interface{}) bool {
				if _, ok := iy.(float64); ok {
					return ix.(float64) > iy.(float64)
				}

				return false
			}
		}

		return nil
	}(anyTypes[0])

	max := anyTypes[0]
	for i := 1; i < len(anyTypes); i++ {
		if isGreatThan(anyTypes[i], max) {
			max = anyTypes[i]
		}
	}

	fmt.Println(max)
}

func Q27_Channel() {
	chValues := make(chan int)
	chQuit := make(chan bool)

	go print(chValues, chQuit)
	for i := 0; i < 10; i++ {
		chValues <- i
	}

	chQuit <- true
}

func print(chValues chan int, chQuit chan bool) {
	for {
		select {
		case c := <-chValues:
			fmt.Println(c)

		case <-chQuit:
			break
		}
	}
}

func Q30_WC() {
	rawString := `I Neko You Neko
	Muhahahah
	I am God of Death`

	numberOfLines := 0
	for _, c := range rawString {
		if c == '\n' {
			numberOfLines++
		}
	}

	if numberOfLines > 0 {
		numberOfLines++
	}

	fmt.Println("number of charaters", len(rawString))
	fmt.Println("number of words:", len(strings.Fields(rawString)))
	fmt.Println("number of lines:", numberOfLines)
}

func Q31_PrintNotDuplicate() {
	input := []byte{'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	result := make([]byte, 0)

	for _, c := range input {
		if isByteInSlice(c, result) == false {
			result = append(result, c)
		}
	}

	fmt.Println(string(result))
}

func isByteInSlice(b byte, byteSlice []byte) bool {
	if len(byteSlice) == 0 {
		return false
	}

	for _, tb := range byteSlice {
		if b == tb {
			return true
		}
	}

	return false
}

func Q32_Quine() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("get working directory error:", err)
		return
	}

	path := dir + "/part2.go"
	content, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("read file from path '", path, "' err: ", err)
		return
	}

	fmt.Println(string(content))
}

func Q33_Echo() {
	fmt.Println("use nc 127.0.0.1 8053 to connect me :D")

	listener, err := net.Listen("tcp", ":8053")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go echo(conn)

	}
}

func echo(conn net.Conn) {
	defer conn.Close()

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("do you know why error?")
		return
	}

	_, err = conn.Write([]byte(line))
	if err != nil {
		fmt.Println("why error at write?", err)
	}
}
