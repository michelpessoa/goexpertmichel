package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	fmt.Println("Status code", req.StatusCode)
	fmt.Println("Status code", req.Status)

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// defer é um comando de garantia que execução após todas as linhas de um programa
	fmt.Println("primeira linha sem defer ")
	fmt.Println("Segunda linha sem defer")
	fmt.Println("Terceira linha sem defer")

	fmt.Println("primeira linha ")
	defer fmt.Println("Segunda linha com defer")
	fmt.Println("Terceira linha ")
}
