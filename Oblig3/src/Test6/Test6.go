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
type yesNo struct {
	Answer string `json:"answer"`
	Forced bool   `json:"forced"`
	Image  string `json:"image"`
}



func Web1(w http.ResponseWriter, r *http.Request)  {
	var sti1 yesNo
	url := "https://yesno.wtf/api"

	result,_ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr :=json.Unmarshal(body, &sti1)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}

	fp := path.Join("templates", "yesNo.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, sti1)
	if err != nil{
		log.Fatal(err)
	}
}
