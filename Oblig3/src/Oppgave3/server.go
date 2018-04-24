package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	go udpClient()
	go tcpClient()

	for {
		time.Sleep(1000)
	}
}

var quoteOfDay = "QOTD: Tell me and I forget. Teach me and I remember. Involve me and I learn."
var quoteOfDay1 = "QOTD: The only true wisdom is in knowing you know nothing."

func udpClient() {

	fmt.Println("Starter UDP server...")
	serverAddr, err := net.ResolveUDPAddr("udp", ":1717")
	checkErr(err)
	conn, err := net.ListenUDP("udp", serverAddr)
	checkErr(err)

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		_, address, err := conn.ReadFromUDP(buf)
		fmt.Println("UDP klient tilkoblet.")
		conn.WriteToUDP([]byte(quoteOfDay+"\n"), address)
		checkErr(err)
	}
}

func tcpClient() {

	fmt.Println("Starter TCP server...")
	ln, err := net.Listen("tcp", ":1717")
	checkErr(err)

	for {
		conn, _ := ln.Accept()
		fmt.Println("TCP klient tilkoblet.")
		conn.Write([]byte(quoteOfDay1 + "\n"))
		conn.Close()
	}
}


func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
