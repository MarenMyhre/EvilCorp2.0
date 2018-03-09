package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"log"
)

func Adder(done chan bool) {
	b, err := ioutil.ReadFile("tempfile.txt.lock")
	//"The system cannot find the file specified" her får vi igen samme feilmeldig som sier at systemet ikke han finne filen
	if err != nil {
		fmt.Print(err)
	}

	inputS := string(b)

	talletS := strings.Fields(inputS)

	tallet1S := string(talletS[0])
	tallet2S := string(talletS[1])

	tallet1, err := strconv.Atoi(tallet1S)
	//
	//Feilmeldingen "Invalid syntax" kommer her fordi tallet1S ikke er et tall
	if err != nil {
		log.Fatal(err)
	}

	tallet2, err := strconv.Atoi(tallet2S)
	//Feilmeldingen "Invalid syntax" kommer opp her fordi tallet2S ikke er et tall
	if err != nil {
		log.Fatal(err)
	}

	sum := tallet1 + tallet2

	sumS := strconv.Itoa(sum)

	sumB := []byte(sumS)

	ioutil.WriteFile("tempfile.txt.lock", sumB, 0777)

	done <- true
}
