package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"fmt"
)
func main() {
	resp, err := http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
	if err != nil {
		log.Fatal(err)
	}

	var apiResp map [string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(apiResp)
}