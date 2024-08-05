package main

import (
	"fmt"
	"time"
)

// Esse exeplo contem duas funcoes "Produtor" e "Cosumidor" 1-N

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
		fmt.Printf("Consumindo1: %d\n", item)
		time.Sleep(1 * time.Second) 
	}
}

func Consumidor2(){
	for item := range produtos { 
		fmt.Printf("Consumindo2: %d\n", item)
		time.Sleep(1 * time.Second) 
	}
}

func Consumidor3(){
	for item := range produtos { 
		fmt.Printf("Consumindo3: %d\n", item)
		time.Sleep(1 * time.Second) 
	}
}

func main() {
	
	
	go Produtor()
	go Consumidor()
	go Consumidor2()
	go Consumidor3()

	time.Sleep(25 * time.Second)
	fmt.Println("Done.")

}
