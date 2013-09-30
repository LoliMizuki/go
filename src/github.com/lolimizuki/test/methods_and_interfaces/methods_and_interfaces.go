package main

import (
	"fmt"
	"math"
	"strconv"
)

type desc_interface interface {
	get_description() string
	give_you_xy(x, y float64)
}

// type and methond
type MyFloat float64

func (f MyFloat) MyPow2() float64 {
	return float64(f * f) // attention here ... =)
	// return f * f // if return type == MyFloat
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// struct and methond
type Vector2 struct {
	x, y float64
}

func (v *Vector2) scale(s float64) {
	v.x = v.x * s
	v.y = v.y * s
}

func (v *Vector2) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// implements desc_interface

func (v *Vector2) get_description() string {
	str_x := strconv.FormatFloat(v.x, 'f', 10, 64)
	str_y := strconv.FormatFloat(v.y, 'f', 10, 64)
	return "value=(" + str_x + ", " + str_y + ")"
}

// error para type
func (v *Vector2) give_you_xy(x, y float64) {
	v.x += x
	v.y += y
}

func main() {
	fmt.Println("-- methods --")
	v := &Vector2{3, -4}
	fmt.Println("v=", v)
	fmt.Println("v.abs()", v.abs())
	v.scale(6)
	fmt.Println("v.scale()", v)

	ff := MyFloat(-math.Sqrt2)
	fmt.Println(ff.Abs())
	fmt.Println(MyFloat(9).MyPow2())

	s := strconv.FormatFloat(12.34567, 'f', 2, 64) //  'e', 'E', 'f', 'g', and 'G'
	fmt.Println("convert float to string", s)

	fmt.Println("-- interface --")
	var a desc_interface
	a = v
	fmt.Println(a.get_description())
	a.give_you_xy(10, 20)
	fmt.Println("add(10,20)", a)
}
