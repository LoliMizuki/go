package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var udpProtocol string = "udp4"

func udpClient() {
	fmt.Println("udp client up")

	udpAddr, err := net.ResolveUDPAddr(udpProtocol, HostAddress)

	conn, err := net.DialUDP(udpProtocol, nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("write nothing"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func udpServer() {
	fmt.Println("udp server up")

	udpServerAddr, err := net.ResolveUDPAddr(udpProtocol, HostAddress)

	conn, err := net.ListenUDP(udpProtocol, udpServerAddr)
	checkError(err)

	for {
		handlerUDPConn(conn)
	}
}

func handlerUDPConn(conn *net.UDPConn) {
	readBuffer := make([]byte, 512)
	_, addr, err := conn.ReadFromUDP(readBuffer[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
