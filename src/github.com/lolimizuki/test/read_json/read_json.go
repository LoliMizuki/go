package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
)

type JsonStruct struct {
	Name string
	Age  int
}

func main() {
	// make_json_string()
	// read_unknow_json_format_from_string()
	// read_unknow_json_format_from_file()
	use_sample_json()
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
	jsonString := `{"Name":"Yuyuko","Age":999}`

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

func use_sample_json() {
	js, err := simplejson.NewJson([]byte(`
	{
		"test": {
		    "array": [1, "2", 3],
		    "int": 10,
		    "float": 5.150,
		    "bignum": 9223372036854775807,
		    "string": "simplejson",
		    "bool": true
		}
	}`))

	if err != nil {
		fmt.Println(err)
		return
	}

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Println(arr)
	fmt.Println(i)
	fmt.Println(ms)
	// 	[1 2 3]
	// 	10
	// 	simplejson

	// read from file

	jsonFile, _ := os.Open("long_json_for_read.json")
	defer jsonFile.Close()

	var b [1024]byte
	n, err := jsonFile.Read(b[:])
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonString := string(b[:n])
	fmt.Println("origin json string", jsonString)

	jsonContent, err := simplejson.NewJson(b[:n])
	if err != nil {
		fmt.Println(err)
		return
	}

	arr2, _ := jsonContent.Array()
	fmt.Println("Len=", len(arr2))
	fmt.Println("Index 0", arr2[0])

	type SimpleStruct struct {
		Boy       string `json:"man"`
		Girl      string `json:"woman, omitempty"`
		Hideyoshi string
	}

	sampleMap := make(map[string]interface{})
	sampleMap["a"] = SimpleStruct{"Rere", "Miz", "Nekoko"}
	sampleMap["b"] = func() []int {
		array := make([]int, 10)
		for i := 0; i < 10; i++ {
			array[i] = i
		}

		return array
	}()

	bytes, err := json.Marshal(sampleMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))
}
