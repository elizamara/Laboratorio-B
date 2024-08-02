package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("./src/templates/*.html"))


//define um tipo para manipular os dados do html
type Data struct {
    Nome           string `json:"nome"`
    Menssagem      string `json:"menssagem"`
}


//verifica se consegue pegar o arquivo e renderiza com execute o dado em /submit depois 
func renderizaTemplate( w http.ResponseWriter, tmpl string, data interface{} ){
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w , err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// renderizar o template da pagina index
func indexHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w,"index.html", nil)
}

// renderizar o template da pagina dinamica 
func dinamicaHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w,"dinamica.html", nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    r.ParseForm()
    data := r.FormValue("data")
    formData := Data{
		Nome: data,
		Menssagem:"É com imensa alegria e entusiasmo que lhe parabenizamos por esta conquista extraordinária! Ganhar na loteria é um feito raro e espetacular, e você acaba de transformar esse sonho em realidade. Que emoção indescritível deve estar sentindo ao saber que a sorte brilhou intensamente para você!",
	}

    // Exibir os dados enviados no servidor
    fmt.Println("Dados recebidos:", formData.Nome)

    // Responder com uma mensagem
    fmt.Fprintf(w, "Você enviou: %s \n", formData.Nome)
	fmt.Fprintf(w, "Querido/a  %s",formData.Nome )
	fmt.Fprintf(w, " \n  %s",formData.Menssagem )
}

func main(){
	
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dinamica", dinamicaHandler)
	http.HandleFunc("/submit", submitHandler)
	
	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}