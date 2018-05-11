package main

import (
	"log"
	"net"
)

var (
	quote = "\"It's nice to be important, but it's more important to be nice'.\"\n"
)

func handleTCP(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte(quote))
	if err != nil {
		log.Printf("Error during typing: %v", err)
	}
}

func QotdTCP() {
	listener, err := net.Listen("tcp", ":17")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error during typing: %v", err)
		}
		go handleTCP(conn)
	}
}

func handleUDP(conn *net.UDPConn) {
	b := make([]byte, 1)
	_, addr, err := conn.ReadFromUDP(b)
	if err != nil {
		log.Printf("Error during typing: %v", err)
	}
	_, err = conn.WriteToUDP([]byte(quote), addr)
	if err != nil {
		log.Printf("Error during typing: %v", err)
	}
}

func QotdUDP() {
	addr, err := net.ResolveUDPAddr("udp", ":17")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		handleUDP(listener)
	}
}

func main() {
	go QotdTCP()
	go QotdUDP()
	log.Println("Qotd listening for connections on port 17.")
	select {

	}
}

