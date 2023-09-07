package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	michel := Cliente{
		Nome:  "Michel",
		Idade: 42,
		Ativo: true,
	}

	minhaEmpresa := Empresa{
		Nome: "Itau",
	}

	Desativacao(michel)

	Desativacao(minhaEmpresa)

	fmt.Printf("Nome da minha empresa %s", minhaEmpresa.Nome)

}
