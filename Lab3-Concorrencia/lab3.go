package main

import (
	"fmt"
)

/* Exemplo simples */

// Código inspirado na apostila fornecida na disciplia
// Modifiquei para executar gorotinas concorrentes
//1 cirei uma função para utilizar canal

func Comunicacao1(canal1 chan int){
	canal1 <- 13
}

func Comunicacao2(palavra chan string){
	palavra <- "Aprovado"
}

func main() {
	canal1:= make(chan int)
	go Comunicacao1(canal1)

	canal2 := make(chan string)
	go Comunicacao2(canal2)

	mostrar:= <- canal1
	fmt.Println(mostrar)
	fmt.Println("-----------")

	mostrar2:= <- canal2
	fmt.Println(mostrar2)
}
