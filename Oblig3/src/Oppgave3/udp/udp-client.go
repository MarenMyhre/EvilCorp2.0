package main

import (
	"fmt"
	"net"
)

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}

func main() {
	ServerAddr,err := net.ResolveUDPAddr("udp","8080")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "8080")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	Conn.Write([]byte(""))
	if err != nil {
		fmt.Println(err)
	}
}