package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os/signal"
	"syscall"
)

func sig() {

	sign := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGINT) //Sender signal ved å bruke Ctrl+C)

	go func() {

		signals := <-sign
		switch signals {
		case syscall.SIGINT:
			fmt.Println("Avslutter...")
			os.Exit(0)

		}
	}()
}


func addtofile() {
	var n1 int
	var n2 int

	fmt.Println("Enter number: ")
	fmt.Scan(&n1)
	fmt.Println("Enter number: ")
	fmt.Scan(&n2)

	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f, "%d\n%d", n1, n2); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func checkErr(e error) {
	if e != nil {
	panic(e)
	}
}

func readResult(path string) {

data, err := ioutil.ReadFile(path)
checkErr(err)

tempData := string(data)
stringData := strings.Split(tempData, " ")
temp := stringData[len(stringData)-2]

result, err := strconv.Atoi(temp)
checkErr(err)

fmt.Println("Resultat: ", result)
}


func main() {
	sig()
	addtofile()
	//sumfromfile()
	readResult("file.txt")
}