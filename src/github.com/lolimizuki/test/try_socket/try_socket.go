package main

import (
	"fmt"
	"net"
	"os"
)

var (
	UseTCPVersion bool   = false
	HostAddress   string = "localhost:8000"
)

func main() {
	server, client := func() (func(), func()) {
		if UseTCPVersion == true {
			return tcpServer, tcpClient
		} else {
			return udpServer, udpClient
		}
		return nil, nil
	}()

	if len(os.Args) > 1 {
		server()
	} else {
		client()
	}

	// parseIP()
}

func parseIP() {
	ipstrs := []string{"163.23.14.5", "aa.bb.cc.dd", "555.666.22.33"}

	for _, ipstr := range ipstrs {
		ip := net.ParseIP(ipstr)
		if ip != nil {
			fmt.Println("You got ip:", ip)
		} else {
			fmt.Println("not ip:", ipstr)
		}
	}
}

func checkError(err error) {
	if err == nil {
		return
	}

	fmt.Println(err)
	os.Exit(1)
}
