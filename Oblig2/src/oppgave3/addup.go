package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go readInput(c)
	time.Sleep(7 * time.Second)
	go addUp(c)
	time.Sleep(7 * time.Second)
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