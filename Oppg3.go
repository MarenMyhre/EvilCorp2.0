package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	sign := make(chan os.Signal, 2)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTSTP) //Sender signal ved å bruke Ctrl+C eller Ctrl+Z

	go func() {

		signals := <-sign
		switch signals {
		case syscall.SIGINT:
			fmt.Println("Goodbye world")
			os.Exit(0)

		case syscall.SIGTSTP:
			fmt.Println("Farewell")
			os.Exit(1)

		}
	}()

	for i := 0; 1 < 100; i++ {
		fmt.Println("Hello world")

	}
}
