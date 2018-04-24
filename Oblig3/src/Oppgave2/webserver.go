package Oppgave2

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
	http.HandleFunc("/2", Web2)
	http.HandleFunc("/3", Web3)
	http.HandleFunc("/4", Web4)
	http.HandleFunc("/5", Web5)
	http.ListenAndServe(":8080", nil)

}
type yesNo struct {
	Answer string `json:"answer"`
	Forced bool   `json:"forced"`
	Image  string `json:"image"`
}
type GoT struct {
	Name        string   `json:"name"`
	Gender      string   `json:"gender"`
	Culture     string   `json:"culture"`
	Born        string   `json:"born"`
	Titles      []string `json:"titles"`
	Aliases     []string `json:"aliases"`
	TvSeries    []string `json:"tvSeries"`
	PlayedBy    []string `json:"playedBy"`
}
type Fox struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}
type Comic struct {
	Num        int    `json:"num"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
}
type Comic2 struct {
	Num        int    `json:"num"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`

}


func Web1(w http.ResponseWriter, r *http.Request) {
	var sti1 yesNo
	url := "https://yesno.wtf/api"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "yesNo.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil {
		log.Fatal(err)
	}
}

func Web2(w http.ResponseWriter, r *http.Request)  {
	var sti2 GoT
	url := "https://anapioficeandfire.com/api/characters/583"

	result,_ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti2)
	if jsonErr != nil{
	log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "GoT.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
	log.Fatal(err)
	}
	err = tmpl.Execute(w, sti2)
	if err != nil{
	log.Fatal(err)
	}
}

func Web3(w http.ResponseWriter, r *http.Request)  {
	var sti3 Fox
	url := "https://randomfox.ca/floof/"

	result,_ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti3)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "Fox.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti3)
	if err != nil{
		log.Fatal(err)
	}
}

func Web4(w http.ResponseWriter, r *http.Request)  {
	var sti4 Comic
	url := "https://xkcd.com/614/info.0.json"

	result,_ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti4)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "Comics.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti4)
	if err != nil{
		log.Fatal(err)
	}
}

func Web5(w http.ResponseWriter, r *http.Request)  {
	var sti5 Comic2
	url := "https://xkcd.com/info.0.json"

	result,_ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &sti5)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "Comics2.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti5)
	if err != nil{
		log.Fatal(err)
	}
}



