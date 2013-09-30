package main

import (
	"fmt"
	"time"
)

const LIMITED = 31

var FibsCache [LIMITED]uint64

func main() {
	TimeProfile(FibWithoutCache)
	TimeProfile(FibWithCache)
}

func TimeProfile(fibFunc func(int) uint64) {
	fmt.Printf("\n")

	start := time.Now()
	for i := 0; i < LIMITED; i++ {
		fmt.Printf("%v, ", fibFunc(i))
	}
	end := time.Now()

	delta := end.Sub(start)

	fmt.Println("take time: ", delta)
}

func FibWithoutCache(n int) uint64 {
	if n <= 1 {
		return 1
	}

	return FibWithCache(n-1) + FibWithCache(n-2)
}

func FibWithCache(n int) uint64 {
	if FibsCache[n] != 0 {
		return FibsCache[n]
	}

	if n <= 1 {
		FibsCache[n] = 1
		return FibsCache[n]
	}

	FibsCache[n] = FibWithCache(n-1) + FibWithCache(n-2)
	return FibsCache[n]
}
