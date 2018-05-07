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
type DayInfo struct {
	Status   int `json:"status"`
	Holidays struct {
		Two0170101 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-01-01"`
		Two0170409 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-04-09"`
		Two0170413 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-04-13"`
		Two0170414 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-04-14"`
		Two0170417 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-04-17"`
		Two0170501 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-05-01"`
		Two0170517 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-05-17"`
		Two0170525 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-05-25"`
		Two0170604 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-06-04"`
		Two0170605 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-06-05"`
		Two0171225 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-12-25"`
		Two0171226 []struct {
			Name     string `json:"name"`
			Date     string `json:"date"`
			Observed string `json:"observed"`
			Public   bool   `json:"public"`
		} `json:"2017-12-26"`
	} `json:"holidays"`
}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 DayInfo
	url := "https://holidayapi.com/v1/holidays?key=ee17c954-a67a-43bf-afb4-9c54ac05ab2b&country=NO&year=2017"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("Templates", "Holi.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}
