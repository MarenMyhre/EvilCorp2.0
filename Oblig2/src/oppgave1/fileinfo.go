package main

import (
	"os"
	"log"
	"fmt"

	"strings"
)

func main() {

	filePath := strings.Join(os.Args[1:], "")
	fi, err :=os.Lstat(filePath)

	if err !=nil {
			log.Fatal(err)
	}

	mode :=fi.Mode()

	fmt.Println("File information: ", fi.Name())
	fmt.Println("Size", fi.Size(), "in bytes")
	fmt.Println(fi.Size() / 1000, "KB")
	fmt.Println(fi.Size() / 1000000, "MB")
	fmt.Println(fi.Size() / 1000000000, "GB")


	switch{
	case mode.IsDir():
		fmt.Println("Is a directory")
	default:
		fmt.Println("Is not a directory")
	}

	switch {
	case mode.IsRegular():
		fmt.Println("Is a regular file")
	default:
		fmt.Println("Is not a regular file")
	}

	switch {
	case mode&os.ModePerm == 0777:
		fmt.Println("Has Unix pernission bits")
	default:
		fmt.Println("Has not Unix pernission bits")
	}

	switch {
	case mode&os.ModeAppend != 0:
		fmt.Println("Is append only")
	default:
		fmt.Println("Is not a device file")
	}

	switch {
	case mode&os.ModeSymlink != 0:
		fmt.Println("In a symbolic link")
	default:
		fmt.Println("Is not a symbolic link")
	}

}

