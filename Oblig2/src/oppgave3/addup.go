package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)
func signaler() {

	sign := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGINT) //Sender signal ved å bruke Ctrl+C)

	go func() {

		signals := <-sign
		switch signals {
		case syscall.SIGINT:
			fmt.Println("Du har valgt å avslutte denne prosessen.")
			os.Exit(0)

		}
	}()
}

func main() {
	signaler()

	c := make(chan int)
	go readInput(c)
	time.Sleep(7 * time.Second)
	go addUp(c)
	time.Sleep(4 * time.Second)
}

func readInput(c chan int) {

	var n1 int
	var n2 int

	fmt.Println("Enter number: ")
	fmt.Scan(&n1)
	fmt.Println("Enter number: ")
	fmt.Scan(&n2)

	c <- n1 
	c <- n2

	res := <-c 
	fmt.Println("Result: ", res)

}

func addUp(c chan int) {

	n1, n2 := <-c, <-c 
	res := (n1 + n2)

	c <- res 

}
