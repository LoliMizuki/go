package qzstack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := stackWithPushOneTwoThree()
	assert((stack.String() == "[0:1] [1:2] [2:3] "), "fail", t)
}

func TestPop(t *testing.T) {
	stack := stackWithPushOneTwoThree()

	assert((stack.Pop() == 3), "pop not 3", t)
	assert((stack.Pop() == 2), "pop not 3", t)
	assert((stack.Pop() == 1), "pop not 3", t)
}

func stackWithPushOneTwoThree() Stack {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	return stack
}

func assert(condition bool, messaage string, t *testing.T) {
	if !condition {
		t.Errorf("test fail")
	}
}
