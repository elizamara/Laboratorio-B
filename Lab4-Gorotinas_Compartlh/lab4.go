package main

import (
	"fmt"
	"time"
)

// Esse exeplo contem duas funcoes "Produtor" e "Cosumidor" utilizando uma espaço
// compartilhado que chamo de produto.

var produtos = make(chan int, 10) 

// Produtor envia dados
func Produtor(){
	for i := 0; i < 20; i++ { 
		fmt.Printf("Enviado: %d\n", i)
		produtos <- i 
		time.Sleep(500 * time.Millisecond) 
	}
	close(produtos)
}

// Retira dados da memoria que seria o produtor 
func Consumidor(){
	for item := range produtos { 
		fmt.Printf("Consumindo: %d\n", item)
		time.Sleep(1 * time.Second) 
	}
}


func main() {
	
	
	go Produtor()
	go Consumidor()

	time.Sleep(25 * time.Second)
	fmt.Println("Done.")

}
