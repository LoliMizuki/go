package error_handler

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// define my error with When and Why
type MizError struct {
	When time.Time
	Why  string
}

// Error()
func (e *MizError) Error() string {
	return fmt.Sprint("at %v, %s", e.When, e.Why)
}

// String()
func (e *MizError) String() string {
	return fmt.Sprint("at %v, %s", e.When, e.Why)
}

// Test Parse with Go idoiom
type ParseError struct {
	index int
	Word  string
	Err   error
}

func MizParse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg %v", r)
			}
		}
	}()

	fileds := strings.Split(input, " ")
	numbers = fileds2int(fileds)
	return
}

func fileds2int(fileds []string) (numbers []int) {
	if len(fileds) == 0 {
		panic("no word to parse")
	}

	for ix, field := range fileds {
		result, err := strconv.Atoi(field)

		if err != nil {
			panic(&ParseError{ix, field, err})
		}

		numbers = append(numbers, result)
	}

	return
}
