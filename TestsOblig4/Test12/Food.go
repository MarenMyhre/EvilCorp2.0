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
	http.HandleFunc("/2", Web1)
	http.ListenAndServe(":8080", nil)
}

type Food struct {
	Title   string  `json:"title"`
	Version float64 `json:"version"`
	Href    string  `json:"href"`
	Results []struct {
		Title       string `json:"title"`
		Href        string `json:"href"`
		Ingredients string `json:"ingredients"`
		Thumbnail   string `json:"thumbnail"`
	} `json:"results"`
}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 Food
	url := "http://www.recipepuppy.com/api/"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("Template", "Food.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}

