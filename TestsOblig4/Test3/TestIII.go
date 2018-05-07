package main

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"path"
	"html/template"
	"net/http"
)
func main() {
	http.HandleFunc("/1", api3)
	http.ListenAndServe(":8080", nil)
}
func api3(w http.ResponseWriter, r * http.Request){

	type Sites struct{
		County string `json:"county"`
		Notifications string `json:"notifications"`
		User_fees string `json:"user_fees"`
		Web_site string `json:"web_site"`
	}

	url := "https://data.mo.gov/resource/b69t-kpq2.json"


	result, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var warningSites []Sites
	jsonErr := json.Unmarshal(body, &warningSites)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fp := path.Join("templates","Strom.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, warningSites); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
