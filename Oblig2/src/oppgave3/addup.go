package main

import (
	"os"
	"strconv"
	"fmt"
	"log"
	"os/signal"
)

func main() {
	d := make(chan os.Signal, 2)
	signal.Notify(d, os.Interrupt,)
	go func() {
		<-d
		fmt.Println("End of process")
		os.Exit(1)
	}()

	functionA()
}

func functionA() {
	input1 := os.Args[1]
	input2 := os.Args[2]

	tallet1, err := strconv.Atoi(input1)
	//Feilmeldingen "Invalid syntax" vises når den første påstanden ikke er et tall
	if err != nil {
		log.Fatal(err)
	}

	tallet2, err := strconv.Atoi(input2)
	//Feilmeldingen "Invalid syntax" vises når den andre påstanden ikke er et tall
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan int)

	go functionB(tallet1, tallet2, c)

	resultat := <-c

	fmt.Println(resultat)
}

func functionB(tallet1 int, tallet2 int, c chan int) {
	resultat := tallet1 + tallet2

	c <- resultat
}