package main

import (
	"fmt"
	"net/rpc"
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
	rpc_tcp_connect()
	// rpc_http_connect()
}

func rpc_http_connect() {
	fmt.Println("call RPC via HTTP")

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	var reply Arith
	err = client.Call("Arith.Multiple", Args{10, 99}, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)

	var quotient Quotient
	err = client.Call("Arith.Divide", Args{12, 5}, &quotient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
}

func rpc_tcp_connect() {
	fmt.Println("call RPC via TCP")

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	var reply Arith
	err = client.Call("Arith.Multiple", Args{10, 99}, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)

	var quotient Quotient
	err = client.Call("Arith.Divide", Args{12, 5}, &quotient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
}
