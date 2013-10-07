package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	// string
	change_character_in_string()
	take_substring()
	loop_in_string()
	number_of_xx_in_string()

	// array and slice
	array_and_slice_make_init()
	cut_last_elemnet_in_arrary()

	use_lable_to_break()
	map_make_init_usage()

	struct_make_init()

	test_value_implements_interface()
	test_interface_classifier()

	recover_function()
	// simple_open_read_file()
	read_write_file_with_slice_buffer()
}

func change_character_in_string() {
	s := "Hello" // immutable
	b := []byte(s)
	b[0] = 'C'

	fmt.Println("origin string is immutable:", s)
	fmt.Println("[]byte is", string(b))
}

func take_substring() {
	str := "nekoko like everyone come from her back"
	fmt.Println(str[:5])
	fmt.Println(str[5:10])
	fmt.Println(str[10:])
}

func loop_in_string() {
	str := "you cannot pass"
	for i := 0; i < len(str); i++ {
		fmt.Printf("[%v] ", string(str[i]))
	}

	fmt.Println("")

	for _, c := range str {
		fmt.Printf("[%v] ", string(c))
	}

	fmt.Println("")
}

func number_of_xx_in_string() {
	str := "Mizuki小鳩鳩"
	fmt.Println("number of byte:", len(str))
	fmt.Println("number of characters:", utf8.RuneCountInString(str)) // use utf8
	// fmt.Println(len([]int(str)))             // []int ... not work?
}

func array_and_slice_make_init() {
	// make
	array := new([10]int)
	slice := make([]int, 10)
	fmt.Println(array)
	fmt.Println(slice)

	// init
	array2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := array2[3:6]
	fmt.Println(array2)
	fmt.Println(slice2)
}

func cut_last_elemnet_in_arrary() {
	slice := make([]int, 5)
	for i := 0; i < 5; i++ {
		slice[i] = i + 1
	}
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}

func use_lable_to_break() {
	found := false
	array2d := [][]int{{1, 3, 5, 7, 9}, {2, 4, 6, 8, 10}}

	target := 8
	var rowLoc int
	var colLoc int
FOUND:
	for rowIndex, row := range array2d {
		for colIndex, col := range row {
			if col == target {
				found = true
				rowLoc = rowIndex
				colLoc = colIndex
				break FOUND
			}
		}
	}

	fmt.Println("Found result: ", found, "at (", rowLoc, ",", colLoc, ")")
}

func map_make_init_usage() {
	myMap := make(map[string]int)
	myMap["a"] = 1
	myMap["b"] = 2

	myMap2 := map[string]int{"nekoko": 1111, "yuyuko": 2222, "pachi": 3333}

	fmt.Println(myMap)
	fmt.Println(myMap2)

	_, isExist := myMap2["aya"]
	fmt.Println("is value of key 'aya' exist?", isExist)

	delete(myMap2, "nekoko")
	fmt.Println("delete 'nekoko':", myMap2)
}

type MyStruct struct {
	number int
	name   string
}

func struct_make_init() {
	mySct := new(MyStruct)
	mySct2 := &MyStruct{13579, "Mizuki"}

	fmt.Println(mySct)
	fmt.Println(mySct2)

	// ** init and make it as factory methods **
}

func (m *MyStruct) GetA() int {
	return m.number
}

func (m *MyStruct) GetB() string {
	return m.name
}

type MyInterface interface {
	GetA() int
	GetB() string
}

func test_value_implements_interface() {
	mySct := new(MyStruct)

	// 轉為 interface type 做測試
	var testInterface interface{}
	testInterface = mySct

	fmt.Println("test my struct implement MyInterface")
	if v, ok := testInterface.(MyInterface); ok {
		fmt.Println("YES, value is", v)
	} else {
		fmt.Println("No he don't")
	}
}

func test_interface_classifier() {
	classifier(123)
	classifier("aaaa")
	classifier(&MyStruct{13579, "Mizuki"})

	var pi float64 = 3.14
	classifier(pi)
}

func classifier(v interface{}) {
	switch v.(type) {
	case int, int64:
		fmt.Println("is int")

	case string:
		fmt.Println("is string")

	case MyStruct, *MyStruct:
		fmt.Println("is MyStruct or *MyStruct")

	default:
		fmt.Println("I don't care your type")
	}
}

func recover_function() {
	defer func() {
		fmt.Println("function done")
		if x := recover(); x != nil {
			fmt.Println("OMG, you panic")
		}
	}()

	panic("PANIC!!!!")
}

func simple_open_read_file() {
	file, err := os.Open("text_file.txt")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer file.Close()

	iReader := bufio.NewReader(file)
	fmt.Println("content in file")
	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println("::", str)
	}
}

func read_write_file_with_slice_buffer() {
	const NUMBER_OF_BUFFER = 512
	var buffer [NUMBER_OF_BUFFER]byte

	file, err := os.Open("text_file.txt")
	if err != nil {
		fmt.Println("can not open file", err)
	}

	for {
		switch numberOfRead, _ := file.Read(buffer[:]); true {
		case numberOfRead < 0:
			return

		case numberOfRead == 0: // EOF
			fmt.Println("... EOF")
			return

		case numberOfRead > 0:
			fmt.Println(string(buffer[:numberOfRead]))
		}
	}

}
