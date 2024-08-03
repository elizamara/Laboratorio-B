package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Exemplo simples */

// Código inspirado na apostila fornecida na disciplia
// Modifiquei para executar gorotinas concorrentes
//1 cirei uma função para utilizar canal

func Produtor(c chan int){
	
	for i:=0; i<3; i++ {
		c <- rand.Intn(100)	
		fmt.Printf("Enviando: %d\n", <-c)	
		time.Sleep(1 * time.Second)
	}

	close(c)
}


func Consumidor(c chan int){
	for num := range c {
		fmt.Printf("Retirando do canal: %d\n", num)
	}
}

func main() {
	c := make(chan int, 3)
	
	go Produtor(c)

	go Consumidor(c)


	time.Sleep(7 * time.Second)

}
