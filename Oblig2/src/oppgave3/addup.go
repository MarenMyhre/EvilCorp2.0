package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var waitGroup sync.WaitGroup
var data chan string

func main() {

	// oppgave 3a
	c := make(chan int)
	go readInput(c)
	time.Sleep(5 * time.Second)
	go addUp(c)
	time.Sleep(5 * time.Second)

}

func print(arg int) {
	fmt.Println("Result from file: ", arg)
}

func readInput(c chan int) {

	var n1 int
	var n2 int

	go func() {
		checkSigint()
	}()

	fmt.Println("Enter num: ")
	fmt.Scan(&n1)
	fmt.Println("Enter num: ")
	fmt.Scan(&n2)

	c <- n1 //sender data via channel
	c <- n2

	res := <-c // mottar resultat fra channel
	fmt.Println("Result: ", res)

}

func addUp(c chan int) {

	n1, n2 := <-c, <-c // mottar data fra readInput()
	res := (n1 + n2)

	c <- res // sender resultat tilbake til readInput()

}

func checkSigint() {

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGHUP,
		syscall.SIGILL)

	s := <-c
	switch s {
	case syscall.SIGINT:
		fmt.Println("Process terminated - SIGINT")
	case syscall.SIGQUIT:
		fmt.Println("Terminal quit - SIGQUIT")
	case syscall.SIGHUP:
		fmt.Println("Hangup - SIGHUP")
	case syscall.SIGILL:
		fmt.Println("Illegal instruction - SIGILL")
	}
}