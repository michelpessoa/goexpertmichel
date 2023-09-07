package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, erro := sum(50, 10)

	if erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
