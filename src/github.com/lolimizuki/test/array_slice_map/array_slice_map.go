package main

import (
	"fmt"
	"sort"
)

func main() {
	// // array
	// my_array := [5]int{1, 2, 3, 4, 5}

	// call_value(my_array)
	// fmt.Println(my_array)

	// call_ref(&my_array)
	// fmt.Println(my_array)

	// slice
	// var my_slice []int
	// 	my_slice := []int{9, 8, 7, 6, 5}
	// 	call_slice(my_slice)

	// do_sort()
	append_apply()
}

func call_value(arr [5]int) {
	arr[0] = 999
}

func call_ref(arr *[5]int) {
	arr[0] = 999
}

func call_slice(slice []int) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func do_sort() {
	var int_slice []int = []int{1, 2, 3, 4, 5, 6, 7}
	pos := sort.SearchInts(int_slice, 5)
	fmt.Println(pos)
}

func append_apply() {
	var int_slice []int
	for i := 0; i < 15; i++ {
		// int_slice[len(int_slice)-1] = i // no work?
		int_slice = append(int_slice, i)
	}
	fmt.Println(int_slice)

	int_slice1 := append(int_slice[:5], int_slice[10:]...)
	fmt.Println("cut 5~10: ", int_slice1)
	// fmt.Println("origin 連動啦: ", int_slice)

	var m map[string]string
}

func funcs_in_map() {
	fmap := map[int] func() {
		1: func() { return 1 },
		2: func() { return 2 },
		3: func() { return 3 },
	}
}
