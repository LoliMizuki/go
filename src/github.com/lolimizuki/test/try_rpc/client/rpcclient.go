package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

var (
	ServerAddress = "localhost:1234"
)

// RPC Structs
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func (q Quotient) String() string {
	return fmt.Sprintf("q=%d, r=%d", q.Quo, q.Rem)
}

type Arith int

func main() {
	// rpc_http_connect()
	// rpc_tcp_connect()
	rpc_json_connect()
}

func rpc_http_connect() {
	fmt.Println("call RPC via HTTP")

	client, err := rpc.DialHTTP("tcp", ServerAddress)
	checkError(err)

	var reply Arith
	err = client.Call("Arith.Multiple", Args{10, 99}, &reply)
	checkError(err)
	fmt.Println(reply)

	var quotient Quotient
	err = client.Call("Arith.Divide", Args{12, 5}, &quotient)
	checkError(err)
	fmt.Println(quotient)
}

func rpc_tcp_connect() {
	fmt.Println("call RPC via TCP")

	client, err := rpc.DialHTTP("tcp", ServerAddress)
	checkError(err)

	var reply Arith
	err = client.Call("Arith.Multiple", Args{10, 99}, &reply)
	checkError(err)
	fmt.Println(reply)

	var quotient Quotient
	err = client.Call("Arith.Divide", Args{12, 5}, &quotient)
	checkError(err)
	fmt.Println(quotient)
}

func rpc_json_connect() {
	fmt.Println("call RPC via JSON")

	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	checkError(err)

	var reply Arith
	err = client.Call("Arith.Multiple", Args{10, 99}, &reply)
	checkError(err)
	fmt.Println(reply)

	var quotient Quotient
	err = client.Call("Arith.Divide", Args{12, 5}, &quotient)
	checkError(err)
	fmt.Println(quotient)
}

func checkError(e error) {
	if e == nil {
		return
	}

	fmt.Println(e)
	os.Exit(1)
}
