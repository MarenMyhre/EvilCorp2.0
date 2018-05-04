package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"encoding/json"
	"log"
	"path"
)

func main() {
	http.HandleFunc("/1", Web1)
	http.ListenAndServe(":8080", nil)
}
type FluAPI []struct {
	UserName    string `json:"user_name"`
	TweetText   string `json:"tweet_text"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	TweetDate   string `json:"tweet_date"`
	Aggravation string `json:"aggravation"`
}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 FluAPI
	url := "http://api.flutrack.org/?s=feverANDcoughORfever"
	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "FluTracker.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}