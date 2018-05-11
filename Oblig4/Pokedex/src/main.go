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
	css := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", css))
	http.HandleFunc("/", Open)
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
		resp = "This Pokémon belongs to generation 1, which is also known as the color generation or the chromatic generation among older fans."
	} else if id > 151 && id < 252 {
		resp = "This Pokémon belongs to generation 2, which can also be referred to as the Gold & Silver series or the metal generation."
	} else if id > 251 && id < 387 {
		resp = "This Pokémon belongs to generation 3, which is sometimes known as the advance or advanced generation."
	} else if id > 386 && id < 494 {
		resp = "This Pokémon belongs to generation 4. This generation focuses on one main region across three games; the Sinnoh region."
	} else if id > 493 && id < 650 {
		resp = "This Pokémon belongs to generation 5. This generation were released in Japan September 2010."
	} else if id > 649 && id < 722 {
		resp = "This Pokémon belongs to generation 6. This generation saw the debut of 72 new Pokémon species!"
	} else if id > 722 && id < 803 {
		resp = "This Pokémon belongs to generation 7, which contains the newest Pokémons. This generation debuted in 2016."
	}else if id < 1 {
		resp = "You have entered an invalid name or ID-number! Please try again. Everything between 1 and 802 are valid numbers."
	}
}
