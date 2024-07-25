package main

import "net/http"

func main(){
// Fornece a pasta onde ta o arqv index.html

	folderSystem := http.FileServer(http.Dir("./src/"))
	http.Handle("/", folderSystem)

	http.ListenAndServe(":8080", nil)
}