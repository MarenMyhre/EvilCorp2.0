package main

import (
	"net/http"
	"io"
)
func main() {
	http.HandleFunc("/", print)
	http.ListenAndServe(":8080", nil)
}

func print(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, client")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}