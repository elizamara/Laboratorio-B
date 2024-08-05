package main

import (
	"fmt"
	"time"
)

// código inspirado na apostila fornecida na disciplia
// Modificado para executar duas funções diferentes "imprimir" e "top" para mostrar o funcionamento
//da Goroutine

func imprimir(n int){
	for i:=0 ; i< 10000; i++ {
		fmt.Printf("%d", n)
		//time.Sleep(200* time.Millisecond)
	}
}

func top(palavra string){
	for i:=0 ; i< 10000; i++ {
		fmt.Printf("%s", palavra)
		time.Sleep(200* time.Millisecond)
	}
}

func main() {
	go top("Rebeca_Andrade")
	top("*EEE*")
	//go imprimir(8)
	//go imprimir(3)
}
