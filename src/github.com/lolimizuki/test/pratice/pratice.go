package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	// complex_test()
	// strings_test()
	// time_test()
	// pointer_test()

	// a, err := strconv.Atoi("wasd")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(a)
	// }

	// fmt.Println(is_equal(12, 13))
	// fmt.Println(is_equal(12, 12))

	// os_open_file()
	// for_test()
	// test_func_type()
	// xyz_caller()

	// x, y, z := get_three_numbers()
	// fmt.Println(x, " ", y, " ", z)
	// give_many(0, 9, 1, 0, 3, 4, 5, 6, 777)
	// type_with_optional_parameters()

	// var value int = 123
	// vType := value.(type)
	// fmt.Println(value, vType)

	// defter_test()
	// func_a()

	f := GetFunc()
	fmt.Println(f(33, 44))

	_, file, line, _ := runtime.Caller(0)
	fmt.Println(file, " ", line)
}

func complex_test() {
	// 複數
	fmt.Printf("%06d\n", 12)
	fmt.Printf("%2.3g\n", 4.22345)

	var c complex64 = 10 + 8i
	fmt.Println(imag(c))
}

func strings_test() {
	// 分割與合併

	str1 := "I am your father!! !!!!!!! <many_space>       </many_space>"

	str_slice1 := strings.Fields(str1)
	for _, v := range str_slice1 {
		fmt.Println("token: ", v)
	}

	var str2 string = "moo][mow][coo][doo][roro"
	str_slice2 := strings.Split(str2, "][")

	for _, v := range str_slice2 {
		fmt.Println("token: ", v)
	}
	fmt.Println(strings.Join(str_slice2, "-join-"))

	// Replace
	str_repace_me := "do you want <w> <w> moo <w> hoo <w>"
	fmt.Println("replace 2: ", strings.Replace(str_repace_me, "<w>", "Yuyuko", 2))
	fmt.Println("replace all: ", strings.Replace(str_repace_me, "<w>", "Yuyuko", -1))
}

func time_test() {
	// time
	t := time.Now()
	fmt.Printf("%04d/%02d/%02d %02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Minute(), t.Second())
}

func pointer_test() {
	i := 5
	fmt.Println("i=", i, " &i=", &i)

	var int_ptr *int
	int_ptr = &i
	fmt.Println("int_ptr=", int_ptr, "*int_ptr=", *int_ptr)
}

func is_equal(x, y int) bool {
	return (x == y)
}

func os_open_file() {
	_, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Oh my God ... what's wrong")
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func for_test() {
	// Exercise 5.5:
	// ver1
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < i+1; j++ {
	// 		fmt.Print("G")
	// 	}
	// 	fmt.Print("\n")
	// }

	// ver2
	curr_str := "G"
	for i := 0; i < 5; i++ {
		fmt.Println(curr_str)
		curr_str += "G"
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("%b, ", i)
	}

	var str string = "貓太郎貓太郎貓太郎貓太郎"
	for i, c := range str {
		fmt.Printf("i= %d: %u\n", i, c)
	}
}

func test_func_type() {
	// 定義一個 func type
	type my_func func(int) string
	var pick_up_func my_func

	rand.Seed(int64(time.Now().Nanosecond()))

	switch choice := rand.Intn(2); choice {
	case 0:
		pick_up_func = test_func1
	case 1:
		pick_up_func = test_func2
	}

	fmt.Println(pick_up_func(199))
}

func test_func1(x int) string {
	return "Yuyuko eat " + strconv.Itoa(x)
}

func test_func2(x int) string {
	return "Pachi read " + strconv.Itoa(x)
}

type XYZ struct {
	x, y, z int
}

func xyz_caller() {
	xyz := XYZ{12, 13, 14}
	fmt.Println(xyz.x)
	change_xyz(&xyz)
	fmt.Println(xyz.x)
}

func change_xyz(xyz *XYZ) {
	xyz.x = 10
	xyz.y = 20
	xyz.z = 30
}

func get_three_numbers() (x, y, z int) {
	x = 10
	y = 20
	z = 30
	return
}

func give_many(attr ...int) {
	for _, a := range attr {
		fmt.Println(a)
	}
}

type MyType struct {
	xxx int
	sss string
}

func type_with_optional_parameters() {
	g := MyType{sss: "aaaa", xxx: 1234}
	fmt.Println(g.sss, " ", g.xxx)
}

func defter_test() string {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	fmt.Println("Func Start")
	fmt.Println("Func Return")
	return "Return!!!"
}

func func_a() {
	y := 333

	a := func() int {
		return 12345 + y
	}()

	fmt.Println(a)
}

func GetFunc() func(int, int) int {
	return func(x, y int) int { return x + y }
}
