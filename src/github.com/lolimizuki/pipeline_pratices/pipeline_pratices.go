package main

import (
	"fmt"
	"sync"
)

func main() {
	// case1_simple()
	case2_merge()
}

func case1_simple() {
	c := gen(2, 3, 4)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func case2_merge() {
	c1 := gen(1, 2, 3)
	c2 := gen(9, 8, 7)

	out := merge(sq(c1), sq(c2))

	for c := range out {
		fmt.Println(c)
	}
}

func gen(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

// 使用了 wait group
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
