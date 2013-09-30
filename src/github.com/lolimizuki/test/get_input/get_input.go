package main

// how to test
// 好像不能在 9o console 中測試
// go build 後, 開啓程式當前的 folder(super+shift+enter), 直接執行程式.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// word_letter_count()
	calculator()
}

// Exercise 12.1: word_letter_count.go
// Write a program which reads text from the keybord. When the user enters ‘S’ in order to signal the end of the input, the program shows 3 numbers:
// i) the number of characters including spaces (but excluding ‘\r’ and ‘\n’)
// ii) the number of words
// iii) the number of lines
func word_letter_count(input string) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Someone want me input: ")
	input, err := inputReader.ReadString('S')

	if err != nil {
		return
	}

	numLine := len(strings.Split(input, "\n"))
	numWords := func(sourceString string) int {
		tokensTemp := strings.Split(sourceString, " ")

		var tokens []string
		for _, s := range tokensTemp {
			s = strings.TrimSpace(s)
			if len(s) > 0 {
				tokens = append(tokens, s)
			}
		}

		return len(tokens)
	}(strings.Replace(input, "\n", " ", -1))

	fmt.Println("number of lines:", numLine)
	fmt.Println("number of words:", numWords)
}

// Exercise 12.2: calculator.go
// Make a simple (reverse polish notation) calculator. This program accepts input from the user in the
// form of integers (maximum 999999) and operators (+, -, *, /).
// The input is like this: number1 ENTER number2 ENTER operator ENTER --> result is displayed.
// The programs stops if the user inputs “q”. Use the package stack you developed in Ex. 11.3

type Stack struct {
	elements []int
}

func (s *Stack) Push(x int) {
	s.elements = append(s.elements, x)
}

func (s *Stack) Pop() int {
	l := len(s.elements)
	if l == 0 {
		return 0
	}

	x := s.elements[l-1]
	s.elements = s.elements[:l-1]

	return x
}

func (s *Stack) String() string {
	return s.String()
}

func (s *Stack) Len() int {
	return len(s.elements)
}

func calculator() {
	var stack Stack

	var input string
	for {
		fmt.Scanln(&input)

		if input == "q" {
			fmt.Println("bye bye")
			break
		}

		if v, err := strconv.Atoi(input); err == nil {
			stack.Push(v)
		} else {
			op := func(input string) func(x, y int) int {
				switch input {
				case "+":
					return func(x, y int) int { return x + y }
				case "-":
					return func(x, y int) int { return x - y }
				case "*":
					return func(x, y int) int { return x * y }
				case "/":
					return func(x, y int) int { return x / y }
				default:
					return nil
				}
			}(input)

			if op != nil {
				fmt.Println(calculatesUntilStackEmpty(&stack, op))
			}
		}
	}
}

func calculatesUntilStackEmpty(stack *Stack, op func(x, y int) int) int {
	if stack.Len() == 0 {
		return 0
	}

	if stack.Len() == 1 {
		return stack.Pop()
	}

	x, y := stack.Pop(), stack.Pop()
	retsult := op(x, y)
	stack.Push(retsult)

	return calculatesUntilStackEmpty(stack, op)
}
