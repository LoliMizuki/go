package error_handler

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMySqrt(t *testing.T) {
	if result, _ := MySqrt(144); result != 12 {
		t.Errorf("sqrt 144 is not 12, is %f", result)
	}
}

func TestError(t *testing.T) {
	_, err := MySqrt(-123)
	if err != nil {
		fmt.Println("Good, err is not nil, err = ", err.Error())
	} else {
		t.Errorf("must have err output")
	}
}

func TestMyDefineError(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	get_value := func() (int32, error) {
		result := rand.Int31() % 100

		if result < 50 {
			return result, &MizError{time.Now(), "result less than 50 error"}
		}

		return result, nil
	}

	for i := 0; i < 10; i++ {
		result, err := get_value()
		if err != nil {
			fmt.Printf("(result=%d)You have my error: %s\n", result, err)
		}
	}
}

func TestFmtErrorf(t *testing.T) {
	getError := func() (int, error) {
		return 123, fmt.Errorf("this error is written by fmt.Errorf(), %s", "lalala")
	}

	_, e := getError()
	fmt.Println(e)
}

func TestRecover(t *testing.T) {
	call_me_get_panic := func() {
		panic("AMAZING PANIC!!! AMAZING PANIC!!!")
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("You have error")
		}
	}()

	fmt.Println("TestRecover Start")
	call_me_get_panic()
	fmt.Println("TestRecover End")
}

func TestMizParse(t *testing.T) {
	var testCases = []string{"12 44 67 88 9", "33 aa bb cc", ""}

	r, e := MizParse(testCases[0])
	if e != nil {
		t.Errorf("e should be nil with test case %s", testCases[0])
	}
	fmt.Println("not error occur, result is ", r)

	r, e = MizParse(testCases[1])
	if e == nil {
		t.Errorf("e should not be nil with test case %s", testCases[1])
	}
	fmt.Println(e)

	r, e = MizParse(testCases[2])
	if e == nil {
		t.Errorf("e should not be nil with test case %s", testCases[2])
	}
	fmt.Println(e)
}
