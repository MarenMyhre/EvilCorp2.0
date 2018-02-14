package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	sign := make(chan os.Signal, 2)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTSTP) //Sender signal ved å bruke Ctrl+C (SIGINT) eller Ctrl+Z (SIGTSTP)

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

// Programmet brukte 27% av CPU og 18 MB av minnet. (Windows)
// Ved å kjøre programmet i GoLand terminalen brukte det rundt 300 MB av minnet.(Windows)
// Programmet brukte 24% av CPU og minnebruken økte så lenge programmet kjørte. Var på ca 2GB etter ett minutt.(Mac) 
