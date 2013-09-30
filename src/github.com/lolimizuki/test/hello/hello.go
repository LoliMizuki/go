package main

import (
	"fmt"
	"github.com/lolimizuki/mypack"
	"math"
	"math/rand"
	"runtime"
	"time"
)

// fail := "FAIL" // := 不可以用在這裡
var success = "Have Nice Day"

type Vertex struct {
	X int
	Y int
}

type number struct {
	num float32
}

type nr number

func add(x int, y int) int {
	return x + y
}

func get_multi_values(x int, y int) (string, int) {
	return "do x*y", x * y
}

// x, y 是回傳值
func get_multi_results(sum int) (x, y int) {
	x = sum / 5 * 2
	y = sum / 5 * 3
	return
}

func declare_vars() {
	var x, y, z int = 1, 2, 3                     // 初始化
	var c, jave, python = false, "suck", true     // 省略 type, 用給值直接決定
	cpp, lisp, ruby := "Landmine", "GOD!!!!", "神" // 使用 := 省略 var 宣告
	fmt.Println(x, y, z, c, jave, python, cpp, lisp, ruby)
}

func for_statement() {
	fmt.Println("for")
	for i := 0; i < 10; i++ {
		fmt.Print(i, "|")
	}

	fmt.Println("")

	// as WHILE
	count := 0
	for count < 10 {
		count++
		fmt.Print("count=", count, "|")
		// if count >= 10 {
		// 	break
		// }
	}

	fmt.Println("")

	// as foreach
	array := []string{"Yuyuko", "Pachi", "Aya"}
	for i, s := range array {
		fmt.Print("#", i, ": ", s, "|")
	}

	fmt.Println("")
}

func test_slice() {
	fmt.Println("slice_test")
	arr := []Vertex{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
	fmt.Printf("length=%d\n", len(arr))
	fmt.Println("full=", arr)

	slice_2_4 := arr[2:4]
	fmt.Println("2:4=", arr[2:4])
	slice_2_4[0].X, slice_2_4[0].Y = slice_2_4[0].X+slice_2_4[0].Y, slice_2_4[0].X*slice_2_4[0].Y
	fmt.Println("after_change", arr)

	// string
	string_array := []string{"阿喵", "小啾", "Nekoko"}
	fmt.Println(string_array)

	// CHECK: copy not work ...
	var copy_string_array []string
	copy(copy_string_array, string_array)
	fmt.Println("-- after copy --")
	fmt.Println(string_array)
	fmt.Println(copy_string_array)
}

func test_slice2() {
	fmt.Println("test slice -- append()")
	// test appent()
	int_array := []int{1, 2, 3, 4, 5, 6}
	int_array = append(int_array, 7)
	int_array = append(int_array, 8, 9, 10)
	fmt.Println(int_array)
}

func test_map() {
	var my_map = map[string]int{
		"top":    480,
		"left":   0,
		"right":  320,
		"botton": 0,
	}
	fmt.Println(my_map)

	var my_map2 = make(map[string]string)
	my_map2["a"] = "A"
	my_map2["b"] = "B"
	my_map2["c"] = "C"
	delete(my_map2, "c")
	delete(my_map2, "d")
	_, ok := my_map2["c"]

	fmt.Println("c in my_map2 exist? ", ok)
	fmt.Println(my_map2)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func test_switch() {
	fmt.Print("OS=")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Opps")
	}

	// no condition swich == clear if ... else if ... else
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	fmt.Println("Rand is ", r)
	switch {
	case 0 <= r && r < 50:
		fmt.Println("0~50")
	case 50 <= r:
		fmt.Println("50~")
	}

	r2 := rand.Intn(5) // 0~4
	switch r2 {
	case 0:
		fallthrough
	case 1:
		fmt.Println("r is 0 or 1")

	case 2:
		fallthrough
	case 3:
		fmt.Println("r is 2 or 3")

	case 4:
		fmt.Println("r is 4")
	}
}

func main() {
	fmt.Println("Hello 世界")

	fmt.Println("Opps", "Give me PI", math.Pi)

	fmt.Println(add(12, 14))

	a, b := get_multi_values(4, 9)
	fmt.Println(a, ": ", b)

	s1, s2 := get_multi_results(10)
	fmt.Println(s1, s2)

	declare_vars()

	fmt.Println(success)

	for_statement()

	p1 := Vertex{Y: 100, X: 12}
	var p2 Vertex = Vertex{33, 44}
	p3 := new(Vertex)
	p3.X, p3.Y = 78, 87

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	test_slice()
	test_slice2()

	test_map()

	fmt.Println(fibonacci(8))
	test_switch()

	fmt.Println("Test My Package")
	fmt.Println(mypack.MyFunction())
	fmt.Println(mypack.MyFunction2())
	fmt.Println(mypack.MyInt)

	aa := number{12.12}
	var bb nr
	// bb = aa
	bb = nr(aa)
	fmt.Println(bb)
}
