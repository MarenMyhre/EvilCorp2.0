package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main()  {
	Response1()
}

type Utsiktpunkt struct {
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

func Response1() {
	//j := `[{"latitude":"58.97349203","name":"Byhaugen","adressenavn":"Carl Sundt Hansens gate 20","longitude":"5.700507617"},{"latitude":"58.89778255","name":"Eikeberget","adressenavn":"Marviksveien 98","longitude":"5.731022334"}]`
	j , err := http.Get("https://hotell.difi.no/api/json/stavanger/utsiktspunkt")
	if err != nil {
		log.Fatal(err)
	}

	defer j.Body.Close()
	contents, err := ioutil.ReadAll(j.Body)
		if err != nil {
			log.Fatal(err)
		}

	xp := []Utsiktpunkt{}

	err = json.Unmarshal(contents, &xp)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("go data: %+v\n", xp)

	for i, v := range xp {
		fmt.Println(i, v)

	}

}