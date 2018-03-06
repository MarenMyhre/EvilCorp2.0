package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"sort"
	"os"
)

var characterCount map[byte]int

func main( ) {
	os.Rename("Oblig2/src/oppgave2/text.txt", "Shakespeare.txt")

	fi, err := ioutil.ReadFile("Shakespeare.txt")
	fileString := string(fi)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Information about Shakespeare.txt")

	fmt.Print("Number of lines in file: ")
	fmt.Println(strings.Count(fileString, "\n"))
	characterCount = make(map[byte]int)
	for i := 0; i < len(fileString); i++ {
		if _, ok := characterCount[fileString[i]]; ok {

		} else {
			characterCount[fileString[i]] = strings.Count(fileString, string(fileString[i]))
		}
	}

	type kv struct {
		Key byte
		Value int
	}

	var ss [ ]kv
	for k, v := range characterCount {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	counter := 0
	for _, kv := range ss {
		if counter < 5 {
			fmt.Printf("%d. Rune: %s, counts: %d\n", counter + 1, string(kv.Key), kv.Value)
		}
		counter++
	}
}