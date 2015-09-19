package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	actionsList := []func(){
		FizzBuzz,
		Strings_1,
		Q6_AvgOfFloat64,
		Q7_RetureCorrent,
		Q9_Stack,
		Q10_ParamArgs,
		Q11_Fib,
		Q15_ReturnPlusFunc,
		Q16_Test,
		Q17_CalculateOfReversePolishNotation,
		Q20_1_DoubleLinkList,
		Q20_2_DoubleLinkList,
		Q24_Interface,
		Q26_Max,
		Q27_Channel,
		Q30_WC,
		Q31_PrintNotDuplicate,
		Q32_Quine,
		Q33_Echo,
	}

	for _, action := range actionsList {
		actionName := runtime.FuncForPC(reflect.ValueOf(action).Pointer()).Name()

		fmt.Println()
		fmt.Println(actionName)
		action()
	}
}
