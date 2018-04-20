package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"path"
	"html/template"
)
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
type Utsiktspunkt struct {
	Entries []struct {
		Latitude    string `json:"latitude"`
		Name        string `json:"name"`
		Adressenavn string `json:"adressenavn"`
		Longitude   string `json:"longitude"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}
type Kommuner struct {
	ID                 string `json:"id"`
	Label              string `json:"label"`
	Lang               string `json:"lang"`
	Contentsummary     string `json:"contentsummary"`
	Owner              string `json:"owner"`
	Manager            string `json:"manager"`
	ContainedItemClass string `json:"containedItemClass"`
	UUID               string `json:"uuid"`
	Containeditems     []struct {
		ID            string        `json:"id"`
		Label         string        `json:"label"`
		Lang          string        `json:"lang"`
		Itemclass     string        `json:"itemclass"`
		UUID          string        `json:"uuid"`
		Status        string        `json:"status"`
		Description   string        `json:"description"`
		Owner         string        `json:"owner"`
		VersionNumber int           `json:"versionNumber"`
		Versions      []interface{} `json:"versions"`
		LastUpdated   string        `json:"lastUpdated"`
		DateSubmitted string        `json:"dateSubmitted"`
		DateAccepted  string        `json:"dateAccepted"`
		Codevalue     string        `json:"codevalue"`
		Narrower      []interface{} `json:"narrower"`
		AlertDate     string        `json:"AlertDate"`
		EffectiveDate string        `json:"EffectiveDate"`
		ValidTo       string        `json:"ValidTo,omitempty"`
		ValidFrom     string        `json:"ValidFrom,omitempty"`
	} `json:"containeditems"`
	ContainedSubRegisters   []interface{} `json:"containedSubRegisters"`
	LastUpdated             string        `json:"lastUpdated"`
	SelectedDOKMunicipality string        `json:"SelectedDOKMunicipality"`
}
type Fylkesnummer struct {
	ID                 string `json:"id"`
	Label              string `json:"label"`
	Lang               string `json:"lang"`
	Contentsummary     string `json:"contentsummary"`
	Owner              string `json:"owner"`
	Manager            string `json:"manager"`
	ContainedItemClass string `json:"containedItemClass"`
	UUID               string `json:"uuid"`
	Containeditems     []struct {
		ID            string        `json:"id"`
		Label         string        `json:"label"`
		Lang          string        `json:"lang"`
		Itemclass     string        `json:"itemclass"`
		UUID          string        `json:"uuid"`
		Status        string        `json:"status"`
		Description   string        `json:"description"`
		Owner         string        `json:"owner"`
		VersionNumber int           `json:"versionNumber"`
		Versions      []interface{} `json:"versions"`
		LastUpdated   string        `json:"lastUpdated"`
		DateSubmitted string        `json:"dateSubmitted"`
		DateAccepted  string        `json:"dateAccepted"`
		Codevalue     string        `json:"codevalue,omitempty"`
		Narrower      []interface{} `json:"narrower"`
		AlertDate     string        `json:"AlertDate"`
		EffectiveDate string        `json:"EffectiveDate"`
		ValidFrom     string        `json:"ValidFrom,omitempty"`
		ValidTo       string        `json:"ValidTo,omitempty"`
	} `json:"containeditems"`
	ContainedSubRegisters   []interface{} `json:"containedSubRegisters"`
	LastUpdated             string        `json:"lastUpdated"`
	SelectedDOKMunicipality string        `json:"SelectedDOKMunicipality"`
}
type NorgeRundt struct {
	Entries []struct {
		TemaVitenskap                      string `json:"tema_vitenskap"`
		TypeKunstOgHaandverk               string `json:"type_kunst_og_haandverk"`
		OffentligeTjenesterOgVeldedighet   string `json:"offentlige_tjenester_og_veldedighet"`
		TypeIdrettOgFysiskAktivitet        string `json:"type_idrett_og_fysisk_aktivitet"`
		TypeDyr                            string `json:"type_dyr"`
		HvaErSpesielt                      string `json:"hva_er_spesielt"`
		Aar                                string `json:"aar"`
		VideoURL                           string `json:"video_url"`
		Hovedperson3Kjonn                  string `json:"hovedperson3_kjonn"`
		TypeLandbrukOgFiske                string `json:"type_landbruk_og_fiske"`
		Hovedperson3Alder                  string `json:"hovedperson3_alder"`
		AntallKvinner                      string `json:"antall_kvinner"`
		TemaKjopOgSalg                     string `json:"tema_kjop_og_salg"`
		TypeMusikere                       string `json:"type_musikere"`
		TypeNaturOgIdrett                  string `json:"type_natur_og_idrett"`
		TemaSamlereEntusiasterOgOppfinnere string `json:"tema_samlere_entusiaster_og_oppfinnere"`
		Hovedperson1Alder                  string `json:"hovedperson1_alder"`
		Tittel                             string `json:"tittel"`
		TypeHistorisk                      string `json:"type_historisk"`
		TypeMat                            string `json:"type_mat"`
		TemaPolitikkOgMedia                string `json:"tema_politikk_og_media"`
		Kommune                            string `json:"kommune"`
		Hovedperson2Kjonn                  string `json:"hovedperson2_kjonn"`
		Hovedperson2Alder                  string `json:"hovedperson2_alder"`
		Tema                               string `json:"tema"`
		Antrekk                            string `json:"antrekk"`
		Hovedperson1Kjonn                  string `json:"hovedperson1_kjonn"`
		TypeByggOgIndustri                 string `json:"type_bygg_og_industri"`
		AntallMenn                         string `json:"antall_menn"`
		TemaNaturOgFriluftsliv             string `json:"tema_natur_og_friluftsliv"`
		Folelse                            string `json:"folelse"`
		Relasjoner                         string `json:"relasjoner"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

func main() {
http.HandleFunc("/1", nettside1)
http.HandleFunc("/2", nettside2)
http.HandleFunc("/3", nettside3)
http.HandleFunc("/4", nettside4)
http.HandleFunc("/5", nettside5)
http.ListenAndServe(":8080", nil)
}

func nettside1(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://hotell.difi.no/api/json/valg/valglokaler/2017")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var valglokaler Valglokaler
	json.Unmarshal(responseData, &valglokaler)

	fp := path.Join("templates", "valglokaler.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, valglokaler); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func nettside2(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://hotell.difi.no/api/json/stavanger/utsiktspunkt")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var utsiktspunkt Utsiktspunkt
	json.Unmarshal(responseData, &utsiktspunkt)

	fp := path.Join("templates", "utsiktspunkt.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, utsiktspunkt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func nettside3(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://hotell.difi.no/api/json/nrk/norge-rundt")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var norgeRundt NorgeRundt
	json.Unmarshal(responseData, &norgeRundt)

	fp := path.Join("templates", "norgeRundt.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, norgeRundt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func nettside4(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://register.geonorge.no/api/subregister/sosi-kodelister/kartverket/kommunenummer-alle.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var kommuner Kommuner
	json.Unmarshal(responseData, &kommuner)

	fp := path.Join("templates", "kommuner.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, kommuner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func nettside5(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://register.geonorge.no/api/subregister/sosi-kodelister/kartverket/fylkesnummer-alle.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var fylkesnummer Fylkesnummer
	json.Unmarshal(responseData, &fylkesnummer)

	fp := path.Join("templates", "fylkesnummer.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, fylkesnummer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}