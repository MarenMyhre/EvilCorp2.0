package main

import (
	"os"
	"testing"
)

func Exists(fil string) bool {
	if _, err := os.Stat(fil); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func TestFil(t *testing.T){
	r := Exists("Dex.html")
	if r == false {
		t.Error("File not found")
	}
}

func TestFil2(t *testing.T){
	r := Exists("Start.html")
	if r == false {
		t.Error("File not found")
	}
}
