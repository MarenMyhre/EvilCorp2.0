package main

import (
	"time"
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
	type WeatherAPI struct {
		Query struct {
			Count   int       `json:"count"`
			Created time.Time `json:"created"`
			Lang    string    `json:"lang"`
			Results struct {
				Channel struct {
					Units struct {
						Distance    string `json:"distance"`
						Pressure    string `json:"pressure"`
						Speed       string `json:"speed"`
						Temperature string `json:"temperature"`
					} `json:"units"`
					Title         string `json:"title"`
					Link          string `json:"link"`
					Description   string `json:"description"`
					Language      string `json:"language"`
					LastBuildDate string `json:"lastBuildDate"`
					TTL           string `json:"ttl"`
					Location struct {
						City    string `json:"city"`
						Country string `json:"country"`
						Region  string `json:"region"`
					} `json:"location"`
					Wind struct {
						Chill     string `json:"chill"`
						Direction string `json:"direction"`
						Speed     string `json:"speed"`
					} `json:"wind"`
					Atmosphere struct {
						Humidity   string `json:"humidity"`
						Pressure   string `json:"pressure"`
						Rising     string `json:"rising"`
						Visibility string `json:"visibility"`
					} `json:"atmosphere"`
					Astronomy struct {
						Sunrise string `json:"sunrise"`
						Sunset  string `json:"sunset"`
					} `json:"astronomy"`
					Image struct {
						Title  string `json:"title"`
						Width  string `json:"width"`
						Height string `json:"height"`
						Link   string `json:"link"`
						URL    string `json:"url"`
					} `json:"image"`
					Item struct {
						Title   string `json:"title"`
						Lat     string `json:"lat"`
						Long    string `json:"long"`
						Link    string `json:"link"`
						PubDate string `json:"pubDate"`
						Condition struct {
							Code string `json:"code"`
							Date string `json:"date"`
							Temp string `json:"temp"`
							Text string `json:"text"`
						} `json:"condition"`
						Forecast []struct {
							Code string `json:"code"`
							Date string `json:"date"`
							Day  string `json:"day"`
							High string `json:"high"`
							Low  string `json:"low"`
							Text string `json:"text"`
						} `json:"forecast"`
						Description string `json:"description"`
						GUID struct {
							IsPermaLink string `json:"isPermaLink"`
						} `json:"guid"`
					} `json:"item"`
				} `json:"channel"`
			} `json:"results"`
		} `json:"query"`
	}

func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 WeatherAPI
	url := "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20weather.forecast%20where%20woeid%20in%20(select%20woeid%20from%20geo.places(1)%20where%20text%3D%22nome%2C%20ak%22)&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "Weather2.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}