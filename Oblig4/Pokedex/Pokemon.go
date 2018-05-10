package main

import (
	"net/http"
	"log"
	"html/template"
	"io/ioutil"
	"encoding/json"
	"path"

)


func main() {
	http.HandleFunc("/1", Open)
	http.HandleFunc("/2", Web)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type Poke struct {
	Forms []struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"forms"`
	Abilities []struct {
			Slot     int  `json:"slot"`
			IsHidden bool `json:"is_hidden"`
			Ability struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"ability"`
		} `json:"abilities"`
	Stats []struct {
			Stat struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"stat"`
			Effort   int `json:"effort"`
			BaseStat int `json:"base_stat"`
		} `json:"stats"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`  //prøv å gang/dele med høyde og vekt for å få mer normale målenheter
	Moves []struct {
			VersionGroupDetails []struct {
				MoveLearnMethod struct {
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"move_learn_method"`
				LevelLearnedAt int `json:"level_learned_at"`
				VersionGroup struct {
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"version_group"`
			} `json:"version_group_details"`
			Move struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"move"`
		} `json:"moves"`
	Sprites struct {
			BackFemale       string `json:"back_female"`
			BackShinyFemale  string `json:"back_shiny_female"`
			BackDefault      string `json:"back_default"`
			FrontFemale      string `json:"front_female"`
			FrontShinyFemale string `json:"front_shiny_female"`
			BackShiny        string `json:"back_shiny"`
			FrontDefault     string `json:"front_default"`
			FrontShiny       string `json:"front_shiny"`
		} `json:"sprites"`
	HeldItems []struct {
			Item struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"item"`
	VersionDetails []struct {
				Version struct {
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"version"`
				Rarity int `json:"rarity"`
			} `json:"version_details"`
		} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Height                 int    `json:"height"`
	IsDefault              bool   `json:"is_default"`
	Species struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"species"`
	ID    int `json:"id"`
	Order int `json:"order"`
	GameIndices []struct {
			Version struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"version"`
			GameIndex int `json:"game_index"`
		} `json:"game_indices"`
	BaseExperience int `json:"base_experience"`
	Types []struct {
			Slot int `json:"slot"`
			Type struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
	Meldinger struct{
		Melding string `string:"Melding"`
	}
}

var info Poke
var url = ""
var resp = ""

func (f *Poke) setResp(response string) {
	f.Meldinger.Melding = response
}

func Open(w http.ResponseWriter, r *http.Request){
	fp := path.Join("Template", "Start.html")
	if r.Method == "GET" {
		t, _ := template.ParseFiles(fp)
		t.Execute(w, nil)
	}else {
		r.ParseForm()
	}
}

func Web(w http.ResponseWriter, r *http.Request) {
	nameID := r.FormValue("ID")
	url= "https://pokeapi.co/api/v2/pokemon/" + string(nameID)
	nameID = ""
	URL := string(url)
	res, err := http.Get(string(URL))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &info)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	info.getResp()
	info.setResp(resp)

	fp2 := path.Join("Template", "Dex.html")
	temp, err := template.ParseFiles(fp2)
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, info)
	if err != nil {
		log.Fatal(err)
	}
}

func (f*Poke) getResp() {
	id := f.ID
	if id < 152 && id > 0 {
		resp = "Dette er generasjon 1"
	} else if id > 151 && id < 252 {
		resp = "Dette er generasjon 2"
	} else if id > 251 && id < 387 {
		resp = "Dette er generasjon 3"
	} else if id > 386 && id < 494 {
		resp = "Dette er generasjon 4"
	} else if id > 493 && id < 650 {
		resp = "Dette er generasjon 5"
	} else if id > 649 && id < 722 {
		resp = "Dette er generasjon 6"
	} else if id > 722 && id < 803 {
		resp = "Dette er generasjon 7"
	}else if id < 1 {
		resp = "Dette er et ugyldig navnet eller ID-nummer. Vær så snill å start programmet på nytt og skrev inn et gyldig navn eller IDnummer. Alt mellom 1-802 er en gyldig ID "
	}
}