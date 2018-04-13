package tcp

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Launching server...")
	fmt.Println("Waiting for client to connect...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", "localhost:8080")
	// accecpt connection on port
	conn, _ := ln.Accept()
	fmt.Println("Connected!")

	for {
		// read and output message from client
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message recieved from client: ", string(message))
		newmessage := message
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
