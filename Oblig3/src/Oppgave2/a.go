package Oppgave2

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

type JsonStruct struct {
	Sted string
}

func main()  {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mate", getMate)
	http.HandleFunc("/1", getParking)
	http.ListenAndServe(":8080", nil)
}

//Client
func foo(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w , r)
}

func getMate(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "mate")
}

func getParking(w http.ResponseWriter, r *http.Request){
	resp, err := http.Get ("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	if err != nil{
		log.Fatal(err)

	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "pre json parre", contents, "\n \n \n")
	data := []JsonStruct{}
	err = json.Unmarshal(contents, data)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Fprint(w, "unmarshalles json", data, "\n \n \n")
	for i, v := range data {
		fmt.Fprint(w, i, "\n \n \n \n ")
		fmt.Fprint(w, v.Sted)
	}
}
