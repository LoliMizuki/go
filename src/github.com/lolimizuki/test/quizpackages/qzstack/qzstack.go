package qzstack

import (
	"fmt"
)

type Stack []int

func NewStack() Stack {
	return Stack(make([]int, 0))
}

func (s *Stack) Push(e int) {
	ints := []int(*s)
	ints = append(ints, e)
	*s = Stack(ints)
}

func (s *Stack) Pop() (e int) {
	ints := []int(*s)
	e = ints[len(ints)-1]
	*s = Stack(ints[:len(ints)-1])

	return
}

func (s Stack) String() string {
	str := ""

	ints := []int(s)
	for ix, iv := range ints {
		str += fmt.Sprintf("[%d:%v] ", ix, iv)
	}

	return str
}
