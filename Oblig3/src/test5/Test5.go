package main

import (
	"net/http"
	"html/template"
	"log"
	"io/ioutil"
	"encoding/json"
	"path"
	)


func main() {
	http.HandleFunc("/1", Web1)
	http.ListenAndServe(":8080", nil)

}
type Valglokaler struct {
	Entries []struct {
		Area              string `json:"area"`
		BoroughName       string `json:"borough_name"`
		AddressLine       string `json:"address_line"`
		PollingPlaceID    string `json:"polling_place_id"`
		BoroughID         string `json:"borough_id"`
		CountyName        string `json:"county_name"`
		PollingPlaceName  string `json:"polling_place_name"`
		GpsCoordinates    string `json:"gps_coordinates"`
		MunicipalityID    string `json:"municipality_id"`
		CountyID          string `json:"county_id"`
		MunicipalityName  string `json:"municipality_name"`
		ElectionDayVoting string `json:"election_day_voting"`
		InfoText          string `json:"info_text"`
		OpeningHours      string `json:"opening_hours"`
		PostalCode        string `json:"postal_code"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

var sti1 Valglokaler

func Web1(w http.ResponseWriter, r *http.Request)  {

	url := "https://hotell.difi.no/api/json/valg/valglokaler/2017"

	result,_ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr :=json.Unmarshal(body, &sti1)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "ValgLokaler.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil{
		log.Fatal(err)
	}
}