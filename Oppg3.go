package main


import "fmt"
import "os"
import "os/signal"
import "syscall"


func main () {

	sign := make(chan os.Signal, 1)
	quitCh = make(chan bool, 1)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; 1 < 100; i++ {
		fmt.Println("Hello world")
	}
case syscall.SIGINT:
fmt.Println("Stop")

case syscall.SIGTERM:
fmt.Println("Terminate")
quitCh <- 0

default:
fmt.Println("Unknown")
quitCh <- 1


code := quitCh
os.Exit(code)
}