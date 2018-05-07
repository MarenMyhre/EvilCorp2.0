package main


import (
	"../Map"
	"net/http"
	"path"
	"log"
	"html/template"
	"io"
)
func main(){
	http.HandleFunc("/1", Web1)
	http.ListenAndServe(":8080", nil)
}

func Web1(w http.ResponseWriter, r *http.Request) {
	hapi := Map.NewV1("91fd5568-8b36-4aaa-a83e-c9ea565686d7") // API key fra Holidays

	holidays, err := hapi.Holidays(map[string]interface{}{
		// Required
		"country": "NO",
		"year":    "2017",
		// Optional
		// "month":    "7",
		// "day":      "4",
		// "previous": "true",
		// "upcoming": "true",
		// "public":   "true",
		// "pretty":   "true",
	})

	if err != nil {
		// Error handling...

		fp := path.Join("Templates", "Holly.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, holidays)
		if err != nil {
			log.Fatal(err)
		}
	}

	io.WriteString(w, "hello, world")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
		//fmt.Println("%#v\n", holidays)
	}
}