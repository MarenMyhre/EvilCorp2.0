package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"

)

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
	addtofile()
	sumfromfile()
	readResult("file.txt")
}