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

	type Sites struct {
		Cod     string  `json:"cod"`
		Message float64 `json:"message"`
		Cnt     int     `json:"cnt"`
		List    []struct {
			Dt   int `json:"dt"`
			Main struct {
				Temp      float64 `json:"temp"`
				TempMin   float64 `json:"temp_min"`
				TempMax   float64 `json:"temp_max"`
				Pressure  float64 `json:"pressure"`
				SeaLevel  float64 `json:"sea_level"`
				GrndLevel float64 `json:"grnd_level"`
				Humidity  int     `json:"humidity"`
				TempKf    float64 `json:"temp_kf"`
			} `json:"main"`
			Weather []struct {
				ID          int    `json:"id"`
				Main        string `json:"main"`
				Description string `json:"description"`
				Icon        string `json:"icon"`
			} `json:"weather"`
			Clouds struct {
				All int `json:"all"`
			} `json:"clouds"`
			Wind struct {
				Speed float64 `json:"speed"`
				Deg   float64 `json:"deg"`
			} `json:"wind"`
			Rain struct {
				ThreeH float64 `json:"3h"`
			} `json:"rain"`
			Snow struct {
			} `json:"snow"`
			Sys struct {
				Pod string `json:"pod"`
			} `json:"sys"`
			DtTxt string `json:"dt_txt"`
		} `json:"list"`
		City struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Coord struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"coord"`
			Country    string `json:"country"`
			Population int    `json:"population"`
		} `json:"city"`
	}
	url := "http://api.openweathermap.org/data/2.5/forecast?q=Oslo%2Cno&units=imperial&appid=0f4ee0e05eebd5458c5e59798b05a962"


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

