package main

import (
	"../Måned"
	"net/http"
)

func main (){
		http.HandleFunc("/1", Måned.Januar)
		http.ListenAndServe(":8080", nil)
}

