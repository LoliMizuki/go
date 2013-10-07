// 沒做, 都是其他雜項的測試 code

package main

import (
	"fmt"
)

type Empty interface{}

type semaphore chan Empty

var NumberOfResources = 10

func main() {
	// sen := make(semaphore, NumberOfResources)
	// shadow()
}

func shadow() {
	x, err := check1()
	if err != nil {
		return
	}

	if y, err := check2(x); err != nil {
		return
	} else {
		fmt.Println(y)
	}
}

func check1() (string, error) {
	return "mimi-live-engine", nil
}

func check2(x string) (string, error) {
	return x + " " + "mk-kill", nil
}

// acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resouces
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// implements mudex

func (s *semaphore) Lock() {
	s.P(1)
}

func (s *semaphore) UnLock() {
	s.V(1)
}
