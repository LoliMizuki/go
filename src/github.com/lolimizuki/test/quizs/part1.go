package main

import (
	"container/list"
	"fmt"
	"github.com/lolimizuki/test/quizpackages/qzstack"
	"strconv"
	"unicode/utf8"
)

func FizzBuzz() {
	for number := 1; number <= 100; number++ {
		outputByte := make([]byte, 0)

		if number%3 == 0 {
			outputByte = append(outputByte, []byte("Fizz")...)
		}

		if number%5 == 0 {
			outputByte = append(outputByte, []byte("Buzz")...)
		}

		if len(outputByte) == 0 {
			outputByte = []byte(strconv.Itoa(number))
		}

		fmt.Printf("%s", string(outputByte))

		if number != 100 {
			fmt.Printf(", ")
		}
	}
}

var testString string = "いんなば abcd 貓貓喵喵"

func Strings_1() {
	fmt.Println(utf8.RuneCountInString(testString))

	runeString := []rune(testString)
	copy(runeString[4:3+len("Mizuki")], []rune("Mizuki"))
	fmt.Println(string(runeString))
}

func Q6_AvgOfFloat64() {
	avgFloat64 := func(floats []float64) float64 {
		sum := 0.0
		for _, f := range floats {
			sum += f
		}

		return sum / float64(len(floats))
	}

	fmt.Println("AVG:", avgFloat64([]float64{12.23, 33.44, 55.66, 100.1}))
}

func Q7_RetureCorrent() {
	retureCorrent := func(x, y int) (rx, ry int) {
		if x > y {
			return y, x
		}
		return x, y
	}

	fmt.Println(retureCorrent(12, 13))
	fmt.Println(retureCorrent(99, 4))
}

type Stack []int

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
		str += fmt.Sprintf("[%d: %v] ", ix, iv)
	}

	return str
}

func Q9_Stack() {
	s := Stack(make([]int, 0))

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s)

	s.Pop()
	s.Pop()
	s.Pop()

	fmt.Println(s)
}

func Q10_ParamArgs() {
	func(ints ...int) {
		for _, iv := range ints {
			fmt.Println(iv)
		}
	}(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
}

func Q11_Fib() {
	fmt.Println(fib(1))
	fmt.Println(fib(4))
	fmt.Println(fib(10))
}

func fib(x int) int {
	if x == 1 || x == 2 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}

func Q15_ReturnPlusFunc() {
	plusX := GetPlusFunc(4)
	fmt.Println(plusX(7))
}

func GetPlusFunc(y int) func(int) int {
	return func(x int) int {
		return x + y
	}
}

func Q16_Test() {
	stack := qzstack.NewStack()
	stack.Push(10)
	stack.Push(101)
	stack.Push(102)

	fmt.Println(stack)
}

func Q17_CalculateOfReversePolishNotation() {
	data1 := []string{"60", "40", "+", "50", "-", "3", "*", "5", "/"}

	RPNCalculator := func(data []string) int {
		numbersInStack := qzstack.NewStack()

		for _, op := range data {
			if n, err := strconv.ParseInt(op, 10, 0); err == nil {
				numbersInStack.Push(int(n))
			} else {
				b := numbersInStack.Pop()
				a := numbersInStack.Pop()

				fmt.Printf("do %d %v %d\n", a, op, b)

				switch op {
				case "+":
					numbersInStack.Push(a + b)
				case "-":
					numbersInStack.Push(a - b)
				case "*":
					numbersInStack.Push(a * b)
				case "/":
					numbersInStack.Push(a / b)
				default:
					panic("unknow op: " + op)
				}
			}
		}

		return numbersInStack.Pop()
	}

	fmt.Println(RPNCalculator(data1))
}

func Q20_1_DoubleLinkList() {
	list := list.New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(4)

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}

	fmt.Println()

	for e := list.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v ", e.Value)
	}
}

type Element struct {
	Value interface{}
	prev  *Element
	next  *Element
}

func (e *Element) Prev() *Element {
	return e.prev
}

func (e *Element) Next() *Element {
	return e.next
}

type List struct {
	header  *Element
	trialer *Element
}

func NewList() *List {
	return new(List)
}

func (l *List) PushBack(v interface{}) {
	e := new(Element)
	e.Value = v

	if l.header == nil {
		l.header = e
	}

	if l.trialer != nil {
		l.trialer.next = e
	}

	e.prev = l.trialer
	e.next = nil

	l.trialer = e
}

func (l *List) Front() *Element {
	return l.header
}

func (l *List) Back() *Element {
	return l.trialer
}

func (l List) String() string {
	s := "["
	for e := l.Front(); e != nil; e = e.Next() {
		s += fmt.Sprintf("%v ", e.Value)
	}

	return s[:len(s)-1] + "]"
}

func Q20_2_DoubleLinkList() {
	list := NewList()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	fmt.Printf("%s\n", list)

	for e := list.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v ", e.Value)
	}
}
