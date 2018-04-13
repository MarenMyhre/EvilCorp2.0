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
	ServerAddr,err := net.ResolveUDPAddr("udp","CHANGEME:27015")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "192.168.1.129:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	Conn.Write([]byte("Mote Fr 5.5 14:45 Flaaklypa"))
	if err != nil {
		fmt.Println(err)
	}
}