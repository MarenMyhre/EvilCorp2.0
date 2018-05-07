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
	http.HandleFunc("/3", Web1)
	http.ListenAndServe(":8080", nil)
}

type DOC struct {
	Success   bool   `json:"success"`
	Remaining int    `json:"remaining"`
	Shuffled  bool   `json:"shuffled"`
	DeckID    string `json:"deck_id"`
}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 DOC
	url := "https://deckofcardsapi.com/api/deck/new/shuffle/?deck_count=1"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("Template", "DOC.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}
