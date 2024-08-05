package main

import (
	"fmt"
	"time"
)

// Esse exeplo contem duas funcoes "Produtor" e "Cosumidor" 1-1

var produtos = make(chan int, 10) 

// Produtor envia dados
func Produtor(){
	for i := 0; i < 3; i++ { 
		fmt.Printf("Produtor1: %d\n", i)
		produtos <- i 
		time.Sleep(200 * time.Millisecond) 
	}
	
}

func Produtor2(){
	for i := 0; i < 10; i++ { 
		fmt.Printf("Produtor2: %d\n", i)
		produtos <- i 
		time.Sleep(100 * time.Millisecond) 
	}
	
}


func Produtor3(){
	for i := 0; i < 20; i++ { 
		fmt.Printf("Produtor3: %d\n", i)
		produtos <- i 
		time.Sleep(50 * time.Millisecond) 
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


func main() {
	
	
	go Produtor()
	go Produtor2()
	go Produtor3()
	go Consumidor()
	// go Consumidor2()
	// go Consumidor3()

	time.Sleep(25 * time.Second)
	fmt.Println("Done.")

}
