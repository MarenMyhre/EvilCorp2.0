package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"path"
	"html/template"
)

func main() {
	http.HandleFunc("/1", Web1)
	http.ListenAndServe(":8080", nil)
}
type Håp struct {
	Status   int `json:"status"`
	Holidays []struct {
		Name     string `json:"name"`
		Date     string `json:"date"`
		Observed string `json:"observed"`
		Public   bool   `json:"public"`
	} `json:"holidays"`
}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 Håp
	url := "https://holidayapi.com/v1/holidays?key=ee17c954-a67a-43bf-afb4-9c54ac05ab2b&country=NO&year=2017&month=05"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("Templates", "San.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}