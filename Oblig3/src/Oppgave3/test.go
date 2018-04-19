package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

// only needed below for sample processing
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hallo"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {

	go func(){

		p := make([]byte, 2048)
		addr := net.UDPAddr{
			Port: 8080,
			IP:   net.ParseIP("127.0.0.1"),
		}
		ser, err := net.ListenUDP("udp", &addr)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
			return
		}
		for {
			_, remoteaddr, err := ser.ReadFromUDP(p)
			fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
			if err != nil {
				fmt.Printf("Some error  %v", err)
				continue
			}
			go sendResponse(ser, remoteaddr)
		}
	}()

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
