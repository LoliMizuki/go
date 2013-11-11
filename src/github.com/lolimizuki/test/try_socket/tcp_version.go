package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"
)

var tcpProtocol string = "tcp4"

func tcpClient() {
	fmt.Println("tcp client up")

	// host := "173.194.72.94:80"
	// host := "www.google.com.tw:80"
	tcpAddr, err := net.ResolveTCPAddr(tcpProtocol, HostAddress)
	checkError(err)

	fmt.Println(tcpAddr)

	conn, err := net.DialTCP(tcpProtocol, nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	date, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(date))

	os.Exit(0)
}

func tcpServer() {
	fmt.Println("tcp server up")

	tcpServerAddr, err := net.ResolveTCPAddr(tcpProtocol, HostAddress)

	listener, err := net.ListenTCP(tcpProtocol, tcpServerAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
		// go handleClientContinueRequest(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	defer func() {
		fmt.Println("timeout ==> disconnect client")
		conn.Close()
	}()

	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

func handleClientContinueRequest(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(time.Minute * 2))
	requestBuffer := make([]byte, 128)
	defer conn.Close()

	for {
		readLength, err := conn.Read(requestBuffer)

		if err != nil {
			continue
		}

		if readLength == 0 {
			break
		} else if string(requestBuffer) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
			conn.Close()
		}

		// clear
		requestBuffer = make([]byte, 128)
	}
}
