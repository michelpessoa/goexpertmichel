package main

import (
	"curso-go/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(matematica.A, 20)
	fmt.Println("Resultado: ", s)

	fmt.Println(uuid.New())
}
