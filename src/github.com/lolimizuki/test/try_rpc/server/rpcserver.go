package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

var (
	ServerAddress = ":1234"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (a *Arith) Multiple(args *Args, reply *Arith) error {
	*reply = Arith(args.A * args.B)
	return nil
}

func (a *Arith) Divide(args *Args, quoti *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quoti.Quo = args.A / args.B
	quoti.Rem = args.A % args.B

	return nil
}

func main() {
	// rpc_http_server()
	// rpc_tcp_server()
	rpc_json_server()
}

func rpc_http_server() {
	fmt.Println("RPC Server(HTTP) boot-on")
	rpc.Register(new(Arith)) // 註冊一個 RPC 服務
	rpc.HandleHTTP()         // 使用 HTTP

	err := http.ListenAndServe(ServerAddress, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func rpc_tcp_server() {
	fmt.Println("RPC Server(TCP) boot-on")
	rpc.Register(new(Arith))

	tcpAddress, err := net.ResolveTCPAddr("tcp", ServerAddress)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddress)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// tcp 的 rpc 必須自行 serve
		go rpc.ServeConn(conn)
	}
}

func rpc_json_server() {
	fmt.Println("RPC Server(JSON) boot-on")

	rpc.Register(new(Arith))

	tcpAddress, err := net.ResolveTCPAddr("tcp", ServerAddress)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddress)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

	net.ListenTCP("tcp", tcpAddress)
}

func checkError(e error) {
	if e == nil {
		return
	}

	fmt.Println(e)
	os.Exit(1)
}
