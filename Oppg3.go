package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	sign := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGINT) //Sender signal ved å bruke Ctrl+C

	go func() {

		<-sign
		fmt.Println("Goodbye world")
		os.Exit(0)
	}()

	for i := 0; 1 < 100; i++ {
		fmt.Println("Hello world")

	}
}
