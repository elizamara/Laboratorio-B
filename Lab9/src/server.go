package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Component struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var (
	components []Component
	mu sync.Mutex
	proxID = 1
)


func main(){

	http.HandleFunc("/", home)
	http.HandleFunc("/add", addComponet)
	//http.HandleFunc("/remove", removeComponet)
	http.HandleFunc("/list", listComponet)

	fmt.Println("Server iniciado na porta :8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request){
	
	fmt.Print(w, "Home page")
}


func addComponet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		http.Error(w, "Método invalido", http.StatusMethodNotAllowed)
		return 
	}

	var component Component
	err := json.NewDecoder(r.Body).Decode(&component)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	mu.Lock()
	component.ID = proxID
	proxID ++
	components = append(components, component)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(component)

}


func listComponet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método invalido", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(components)
}