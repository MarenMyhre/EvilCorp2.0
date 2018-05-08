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

type Nameday struct {
	Data struct {
		NameUs string `json:"name_us"`
		NameCz string `json:"name_cz"`
		NameSk string `json:"name_sk"`
		NamePl string `json:"name_pl"`
		NameFr string `json:"name_fr"`
		NameHu string `json:"name_hu"`
		NameHr string `json:"name_hr"`
		NameSe string `json:"name_se"`
		NameAt string `json:"name_at"`
		NameIt string `json:"name_it"`
		NameEs string `json:"name_es"`
		Day    int    `json:"day"`
		Month  int    `json:"month"`
	} `json:"data"`
}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 Nameday
	url := "https://api.abalin.net/get/today"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("Templates", "name.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}
