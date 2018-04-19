package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/1", utsiktspunkt)
	http.HandleFunc("/2", fylkesnummer)
	http.HandleFunc("/3", kommuner)
	http.HandleFunc("/4", norgeRundt)
	http.HandleFunc("/5", valglokaler)
	http.ListenAndServe(":8080", nil)

}

func utsiktspunkt(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "utsiktspunkt.gohtml", nil)
}

func fylkesnummer(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "fylkesnummer.gohtml", nil)
}

func kommuner(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "kommuner.gohtml", nil)
}

func norgeRundt(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "norgeRundt.gohtml", nil)
}

func valglokaler(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "valglokaler.gohtml", nil)
}
