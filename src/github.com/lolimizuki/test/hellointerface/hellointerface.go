package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float64
}

// Square
type Square struct {
	side float64
}

func (s *Square) Area() float64 { // pointer
	return s.side * s.side
}

// Recangle
type Recangle struct {
	Width, Height float64
}

func (r Recangle) Area() float64 { // non-pointer
	return r.Width * r.Height
}

// Circle
type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

type Stringer interface {
	String() string
}

func main() {
	basic_try()
	try_to_identify_interface_value_type()
	try_pointer_value_call_methods()

	// 帶入 sort 函數
	// try_sort_MyData()
}

func basic_try() {
	var shapers []Shaper
	shapers = append(shapers, &Square{12.88})   // pointer
	shapers = append(shapers, Recangle{10, 20}) // it's a value
	shapers = append(shapers, &Circle{6})       // pointer

	for i, _ := range shapers {
		showArea(shapers[i])
	}
}

func showArea(shaper Shaper) {
	// interface switch by type(test interface imper's type)
	switch t := shaper.(type) {
	// pointer
	case *Square:
		fmt.Println("Good Square pointer... ")
	case *Circle:
		fmt.Println("Nice Circle pointer... ")
	case *Recangle:
		fmt.Println("Suck Recangle pointer")

	case Recangle: // value
		fmt.Println("Suck Recangle value")

	default:
		fmt.Println("unknow type %T", t)
	}

	// call Area()
	fmt.Println(shaper.Area())
}

func try_to_identify_interface_value_type() {
	// test value(interface value) has implement interface?
	v := &Square{33.77}
	var vInt Shaper
	vInt = v
	if sv, ok := vInt.(Shaper); ok {
		fmt.Printf("YES, imp the Shaper, Area()=%f", sv.Area())
	}

	fmt.Println("")
}

// Methods set with interface 的注意事項
// * Pointer methods 可以被 Pointer 呼叫
// * Value methods 可以被 Value 呼叫
// * Value-Receiver 可以被 Pointer-Value 呼叫, 因為會先 dereferenced
// * Pointer-Receiver 不可以被 Value 呼叫, 因為因為存在 interface-value 中的 value 沒有其 address
type TestInterface interface {
	CallByValue()
	CallByPointer()
}

type TestStruct struct {
}

func (ts TestStruct) RecvIsValue() {
	fmt.Println("RecvIsValue")
}

func (ts *TestStruct) RecvIsPointer() {
	fmt.Println("RecvIsPointer")
}

func (ts TestStruct) CallByValue() {
	fmt.Println("Call by value")
}

func (ts *TestStruct) CallByPointer() {
	fmt.Println("Call by pointer")
}

func try_pointer_value_call_methods() {
	// 以下皆 ok
	fmt.Println("Value: ")
	ts := TestStruct{}
	ts.RecvIsValue()
	ts.RecvIsPointer()

	fmt.Println("Pointer: ")
	p_ts := &TestStruct{}
	p_ts.RecvIsValue()
	p_ts.RecvIsPointer()

	// 使用 varI 來呼叫
	var varI TestInterface
	// varI = ts // ... (no)
	// varI = &ts // ... (ok)
	varI = p_ts
	varI.CallByValue()
	varI.CallByPointer()
}

// 應用 Compare Function as Sort
// type 要提供"比較"和"交換"兩種實作
type Comparee interface {
	Less(another Comparee) bool
}

type MyData struct {
	number int
	id     string
}

func (self *MyData) Less(another Comparee) bool {
	if anotherMyData, ok := another.(*MyData); ok {
		return self.number < anotherMyData.number
	}
	return false
}

func try_sort_MyData() {
	var comparees []Comparee
	comparees = append(comparees, &MyData{5, "five"})
	comparees = append(comparees, &MyData{3, "three"})
	comparees = append(comparees, &MyData{1, "one"})
	comparees = append(comparees, &MyData{2, "two"})
	comparees = append(comparees, &MyData{4, "four"})

	// DoSortWithoutLessFunc(comparees)

	// var lessFunc func(x, y Comparee) bool = func(x, y Comparee) bool {
	// 	myData_x, _ := x.(*MyData)
	// 	myData_y, _ := y.(*MyData)

	// 	return myData_x.number < myData_y.number
	// }
	// DoSortWithLessFunc(comparees, leeFunc)

	DoSortWithLessFunc(comparees,
		func(x, y Comparee) bool { // 直接 define lessFunc
			myData_x, _ := x.(*MyData)
			myData_y, _ := y.(*MyData)

			return myData_x.number < myData_y.number
		})

	for _, c := range comparees {
		if o, ok := c.(*MyData); ok {
			fmt.Println(o.number, ":", o.id)
		}
	}
}

func DoSortWithoutLessFunc(comparees []Comparee) {
	for pass := 1; pass < len(comparees); pass++ {
		for i := 0; i < len(comparees)-pass; i++ {
			if comparees[i+1].Less(comparees[i]) == false {
				comparees[i+1], comparees[i] = comparees[i], comparees[i+1]
			}
		}
	}
}

func DoSortWithLessFunc(comparees []Comparee, lessFunc func(c1, c2 Comparee) bool) {
	for pass := 1; pass < len(comparees); pass++ {
		for i := 0; i < len(comparees)-pass; i++ {
			if lessFunc(comparees[i], comparees[i+1]) {
				comparees[i+1], comparees[i] = comparees[i], comparees[i+1]
			}
		}
	}
}
