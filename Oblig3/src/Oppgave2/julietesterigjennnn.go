package main

import (
"net/http"
	"net/url"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}


func foo() {
	u := &url.URL{}
	err := u.UnmarshalBinary([]byte("https://register.geonorge.no/api/subregister/sosi-kodelister/kartverket/kommunenummer-alle.json"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", u)
}