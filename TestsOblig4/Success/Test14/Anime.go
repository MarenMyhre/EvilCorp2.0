
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

type Film[]struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Director    string   `json:"director"`
	Producer    string   `json:"producer"`
	ReleaseDate string   `json:"release_date"`
	RtScore     string   `json:"rt_score"`
	People      []string `json:"people"`
	Species     []string `json:"species"`
	Locations   []string `json:"locations"`
	Vehicles    []string `json:"vehicles"`
	URL         string   `json:"url"`
}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 Film
	url := "https://ghibliapi.herokuapp.com/films/"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("Template", "Anime.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}


