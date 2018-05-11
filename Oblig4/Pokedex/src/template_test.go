package main_test

import (
	"os"
	"testing"
	"path"
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
	fp := path.Join("Template", "Start.html")
	r := Exists(fp)
	if r == false {
		t.Error("File not found")
	}
}

func TestFil2(t *testing.T){
	fp1 := path.Join("Template", "Dex.html")
	r := Exists(fp1)
	if r == false {
		t.Error("File not found")
	}
}
