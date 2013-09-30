package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	F1, F2, F3 string
}

func (m MyStruct) String() string {
	return m.F1 + ", " + m.F2 + ", " + m.F3
}

func main() {
	// reflect_with_float()
	reflect_with_struct()
}

func reflect_with_float() {
	var ff float64 = 3.876
	fmt.Println("Type=", reflect.TypeOf(ff), ", Value=", reflect.ValueOf(ff))

	v := reflect.ValueOf(ff)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	fmt.Println(v.Float())
	fmt.Println(v.Interface())

	fmt.Println("CanSet? ", v.CanSet()) // no address, can not set

	p_v := reflect.ValueOf(&ff)           // get from &
	fmt.Println("CanSet? ", p_v.CanSet()) // no address, can not set
	e := p_v.Elem()
	e.SetFloat(66.7788)
	fmt.Println("CanSet? ", e.CanSet(), "Value is ", e.Float())
}

func reflect_with_struct() {
	myStruct := MyStruct{"aa", "bb", "cc"}
	ref := reflect.ValueOf(myStruct)
	fmt.Println("number of field(s)=", ref.NumField())
	fmt.Println("number of method(s)=", ref.NumMethod())
	for i := 0; i < ref.NumField(); i++ {
		fmt.Println(ref.Field(i))
	}
	fmt.Println(ref.Method(0).Call(nil))

	myStruct2 := &MyStruct{"aa", "bb", "cc"}
	elem := reflect.ValueOf(myStruct2).Elem()
	elem.Field(0).SetString("Sachiko")
	fmt.Println(elem.Field(0), elem.Field(1), elem.Field(2))
	fmt.Println(myStruct2)
}
