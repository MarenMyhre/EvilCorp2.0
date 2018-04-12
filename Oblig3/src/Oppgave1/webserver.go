package Oppgave1

import (
	"fmt"
	"net"
	"bufio"
)

func main() {

	fmt.Println("Launching server...")
	fmt.Println("Waiting for client to connect...")

	ln, _ := net.Listen("tcp", "localhost:8081")
	conn, _ := ln.Accept()
	fmt.Println("Connected!")

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message recieved from client: ", string(message))
		newmessage := message
		conn.Write([]byte(newmessage + "\n"))
	}
}
