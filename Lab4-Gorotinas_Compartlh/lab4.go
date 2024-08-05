package main

import (
	"fmt"
	"time"
)

/* Exemplo simples */

// Código inspirado na apostila fornecida na disciplia
// Modifiquei para executar gorotinas concorrentes
//1 cirei uma função para utilizar canal

func Produtor(c chan <- int){
	c <- 13
	time.Sleep(1 * time.Second)
	fmt.Printf("Enviando: %d\n", c)	
}


func Consumidor(c <- chan int){

}


func extra(c chan int){
	for {
		 valor, ok := <-c 
		 if ok {
			fmt.Println(valor)
		 }else{
			break	
		 }
	}

	close(c)
}


func main() {
	c := make(chan int)
	
	go Produtor(c)
	//go Consumidor(c)

}
