package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"nu"`
	Saldo  float64 `json:"s"`
	Nome   string  `json:"n" validate:"gt=0"`
}

type ContaJson struct {
	Numero int     `json:"numero"`
	Saldo  float64 `json:"saldo" validate:"gt=0"`
	Nome   string  `json:"nome"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 12.2, Nome: "Michel"}
	objJson, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(objJson))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"numero":2,"saldo":20.2,"nome":"João Carlos"}`)
	var contaB ContaJson
	err = json.Unmarshal(jsonPuro, &contaB)
	if err != nil {
		panic(err)
	}
	fmt.Println(contaB.Saldo)
	fmt.Println(contaB.Nome)

	jsonPuro2 := []byte(`{"nu":3,"s":30.3,"n":"Carlo Roberto"}`)
	var contaC Conta
	err = json.Unmarshal(jsonPuro2, &contaC)
	if err != nil {
		panic(err)
	}
	fmt.Println(contaC.Saldo)
	fmt.Println(contaC.Nome)

}
