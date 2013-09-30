package mimisort

import (
	"fmt"
	"strconv"
	"testing"
)

// Ints Array
type IntArray []int

func (integers IntArray) Len() int {
	return len(integers)
}

func (integers IntArray) Less(i, j int) bool {
	return integers[i] < integers[j]
}

func (integers IntArray) Swap(i, j int) {
	integers[i], integers[j] = integers[j], integers[i]
}

// stringArray
type StringArray []string

func (strings StringArray) Len() int {
	return len(strings)
}

func (strings StringArray) Less(i, j int) bool {
	return strings[i] < strings[j]
}

func (strings StringArray) Swap(i, j int) {
	strings[i], strings[j] = strings[j], strings[i]
}

// structs array
type MyStruct struct {
	id  string
	num int
}

type MyStructArray []MyStruct

func (myStructs MyStructArray) Len() int {
	return len(myStructs)
}

func (myStructs MyStructArray) Less(i, j int) bool {
	return myStructs[i].id < myStructs[j].id
}

func (myStructs MyStructArray) Swap(i, j int) {
	myStructs[i], myStructs[j] = myStructs[j], myStructs[i]
}

func (m *MyStruct) String() string {
	return "{id=" + m.id + ", num=" + strconv.Itoa(m.num) + "}"
}

func TestIntsArray(t *testing.T) {
	intArr := IntArray{12, 23, 55, 11, 57, 0, 8, 44, 3}
	Sort(intArr)

	if IsSorted(intArr) == false {
		t.Errorf("sort ints array fail with order %v", intArr)
	}
}

func TestStringsArray(t *testing.T) {
	stringsArray := StringArray{"fdsa", "grb", "yttc", "dsdd", "zzzt", "Mizuki", "Pachi", "Totori", "Nekoko", "nekoko", "nakoko"}
	Sort(stringsArray)

	if IsSorted(stringsArray) == false {
		t.Errorf("sort fail in strings array, order: , %v", stringsArray)
	}
}

func TestStructsArray(t *testing.T) {
	myStructs := MyStructArray{
		MyStruct{"Yuyuko", 69696},
		MyStruct{"Nekoko", 123},
		MyStruct{"Pachi", 9090},
		MyStruct{"Apple", 776}}
	Sort(myStructs)

	if IsSorted(myStructs) == false {
		t.Errorf("sort fail in structs array, order: , %v", myStructs)
	}
}
