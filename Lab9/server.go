package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("./src/templates/*.html"))


func indexHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w,"index.html", nil)
}


func dinamicaHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w,"dinamica.html", nil)
}


func main(){

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dinamica", dinamicaHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}