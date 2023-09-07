// package main

// import "fmt"

// func main() {

// 	// Closures - função embutidas dentro do código. (rotinas)
// 	// func() {
// 	// }()

// 	total := func() int {
// 		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) * 2
// 	}()

// 	fmt.Println(total)
// }

//	func sum(numeros ...int) int {
//		total := 0
//		for _, numeros := range numeros {
//			total += numeros
//		}
//		return total
//	}
package main

import "fmt"

// GO não é uma linguagem OO

// STRUCT

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	michel := Cliente{
		Nome:  "Michel",
		Idade: 42,
		Ativo: true,
	}

	fmt.Printf("O nome do cliente é: %s, com a idade: %d e status: %t ", michel.Nome, michel.Idade, michel.Ativo)

	michel.Ativo = false

	fmt.Printf("\nAgora o cliente %s está com o status: %t", michel.Nome, michel.Ativo)
}
