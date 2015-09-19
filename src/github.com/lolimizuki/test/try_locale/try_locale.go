// 只是做個基本的 map 對應, 並沒有真的 locale 化啦 >..<

package main

import (
	"errors"
	"fmt"
	"time"
)

type FuncGivenError func(a, b int) error

func main() {
	// message()
	// timezone()
	TestWrapError()
}

func message() {
	loc := "tw"
	// loc := "en"

	mimiloc := make(map[string]string, 2)
	mimiloc["en"] = "mimi"
	mimiloc["tw"] = "蜜蜜"

	stmtloc := make(map[string]string, 2)
	stmtloc["en"] = "Hello, my sweet %s"
	stmtloc["tw"] = "%s出現了, 您要[攻擊][防禦][逃走][X了她]"

	fmt.Printf(stmtloc[loc], mimiloc[loc])
}

func timezone() {
	timeloc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err)
		return
	}

	now := time.Now()
	now = now.In(timeloc)

	fmt.Println(now.Format(time.RFC3339))
}

func TestWrapError() {
	WrapError(error_given1, 12, 14)
	WrapError(error_given1, 0, 1)
}

func WrapError(funcGenError FuncGivenError, a, b int) {
	if err := funcGenError(a, b); err != nil {
		fmt.Println("error:", err)
	}
}

func error_given1(a, b int) error {
	if a == 0 {
		return errors.New("a is zero")
	}

	fmt.Println(a + b)
	return nil
}
