package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":1717")
	checkErr(err)

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)
	conn.Close()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}