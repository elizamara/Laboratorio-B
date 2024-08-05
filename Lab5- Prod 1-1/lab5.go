package main

import (
	"fmt"
	"time"
)

// Esse exeplo contem duas funcoes "Produtor" e "Cosumidor" 1-1

var produto = make(chan int, 10) 

// Produtor envia dados
func Produtor(){
	produto <- 2
	fmt.Printf("Enviado: %d\n", <-produto) 
	time.Sleep(500 * time.Millisecond) 
	close(produto)
}

// Retira dados da memoria que seria o produtor 
func Consumidor(){
 
	fmt.Printf("Consumindo: %d\n", <-produto)
	time.Sleep(1 * time.Second) 
	
}


func main() {
	
	
	go Produtor()
	go Consumidor()

	time.Sleep(25 * time.Second)
	fmt.Println("Done.")

}
