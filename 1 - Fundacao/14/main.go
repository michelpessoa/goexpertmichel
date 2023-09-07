package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c Cliente) andou() {
	c.nome = "Ana Clara"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

// Espécie de um construtor de uma struct
// func NewConta() *Conta {
// 	return &Conta{saldo: 0}
// }

func (co *Conta) simular(valor int) int {
	co.saldo += valor
	println(co.saldo)
	return co.saldo

}

func main() {
	// michel := Cliente{
	// 	nome: "Michel",
	// }

	// michel.andou()
	// fmt.Printf("O valor da struct com nome %v\n\n\n\n", michel.nome)

	conta := Conta{saldo: 100}

	conta.simular(200)
	//fmt.Printf("Saldo da conta %v", conta.saldo)
	println(conta.saldo)

}
