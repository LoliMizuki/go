package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonStruct struct {
	Name string
	Age  int
}

func main() {
	// make_json_string()
	// read_unknow_json_format_from_string()
	read_unknow_json_format_from_file()
}

func make_json_string() {
	jStruct := JsonStruct{"Yuyuko", 999}
	jsonString, e := json.Marshal(jStruct)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%s", jsonString)
}

func read_unknow_json_format_from_string() {
	var emptyInterface interface{}
	var originStruct JsonStruct
	jsonString := "{\"Name\":\"Yuyuko\",\"Age\":999}"

	json.Unmarshal([]byte(jsonString), &emptyInterface)
	json.Unmarshal([]byte(jsonString), &originStruct)

	fmt.Println(emptyInterface) // 變成一個 map
	fmt.Println(originStruct)
}

func read_unknow_json_format_from_file() {
	file, _ := os.Open("json_for_read.json")
	defer file.Close()

	var bytes [1024]byte
	n, _ := file.Read(bytes[:])

	jsonString := string(bytes[:n])
	fmt.Println("json string is \"", jsonString, "\"")

	var m interface{}
	json.Unmarshal([]byte(jsonString), &m)
	fmt.Println(m)
}
