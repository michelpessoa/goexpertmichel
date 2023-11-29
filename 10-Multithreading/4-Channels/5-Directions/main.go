package main

import "fmt"

// Devido ao posicinamento à direita da "seta" na declaração,
// determina a atribuição do valor na variável
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// Já quando a "seta" está posicionada à esquerda da variável,
// o funcionamento é a leitura e esvaziamento do valor anteriormente na variável
func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
