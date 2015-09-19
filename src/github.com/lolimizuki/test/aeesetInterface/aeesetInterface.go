package main

import (
	"encoding/json"
	"fmt"
)

// var jsonString = `[
//         {"Name": "Platypus", "Order": "Monotremata"},
//         {"Name": "Quoll",    "Order": "Dasyuromorphia"}
// ]`

var jsonString = `{
    "stringsValue": ["aa", "bb", "cc"],
    "strValue": "I am String"
}`

var jsonBytes = []byte(jsonString)

func main() {
	// type Animal struct {
	// 	Name  string
	// 	Order string
	// }

	// var animals []Animal

	// if err := json.Unmarshal(jsonBytes, &animals); err != nil {
	// 	fmt.Println("fail", err)
	// 	return
	// }

	mapFromJson := make(map[string]interface{})
	if err := json.Unmarshal(jsonBytes, &mapFromJson); err != nil {
		fmt.Println("fail", err)
		return
	}

	// for _, a := range animals {
	// 	fmt.Println("Name:", a.Name, "Order:", a.Order)
	// }

	for _, innerLnterface := range mapFromJson {
		AssertInterface(innerLnterface)
	}
}

func AssertInterface(unknowInterface interface{}) {
	switch unknowInterface.(type) {
	case string:
		fmt.Println("string:", unknowInterface.(string))

	case []string:
		fmt.Println("array of strings")

		case ([]interface):
		    fmt.Println("array of something")

	case int:
		fmt.Println("int: ", unknowInterface.(int))

	default:
		fmt.Println("I do't know who you are?")

	}
}
