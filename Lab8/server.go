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


func main(){

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}