package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go readInput(c)
	time.Sleep(5 * 1e9)
	go addUp(c)
	time.Sleep(5 * 1e9)
}

func readInput(c chan int) {

	var n1 int
	var n2 int

	fmt.Println("Enter number: ")
	fmt.Scan(&n1)
	fmt.Println("Enter number: ")
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