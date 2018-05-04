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

type Håp struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt int `json:"dt"`
}
type Main struct {
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
		Country string `json:"country"`
	} `json:"city"`
}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 Håp
	url := "http://api.openweathermap.org/data/2.5/forecast?id=3149318&APPID=bea05aad499f5547900e6ab4cc019d4f"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("Template", "Kr.sand.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}
