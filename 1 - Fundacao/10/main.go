package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado com o valor %t", c.Nome, c.Ativo)
}

func main() {

	michel := Cliente{
		Nome:  "Michel",
		Idade: 42,
		Ativo: true,
	}

	michel.Cidade = "Goiânia"

	fmt.Printf("\n\nO nome do cliente é: %s, com a idade: %d e status: %t ", michel.Nome, michel.Idade, michel.Ativo)

	// michel.Ativo = false

	fmt.Printf("\nO cliente %s, mora na cidade de %s", michel.Nome, michel.Cidade)

	michel.Desativar()

	fmt.Printf("\nAgora o cliente %s está com o status: %t\n", michel.Nome, michel.Ativo)
}
