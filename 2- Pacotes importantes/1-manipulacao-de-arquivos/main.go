package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Olá Mundo!")
	tamanho, err := f.Write([]byte("Escrevendo arquivo em bytes!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)

	f.Close()

	// LEITURA DE ARQUIVOS

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("\n" + string(arquivo) + "\n")

	// LEITURA ARQUIVO POUCO EM POUCO

	arquivo2, err := os.Open("lista.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 100)
	local_time := time.Now()
	fmt.Println("Início do processo de bufferização do arrquivo ", local_time.Format("02/01/2006 00:00:00"))
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		//time.Sleep(3)
		fmt.Println(string(buffer[:n]))
	}
	local_time = time.Now()
	fmt.Println("Termino do processo de bufferização do arrquivo ", local_time.Format("02/01/2006 15:04:05")) // Data e hora obrigatóriamente 02/01/2006 = DD/MM/YYYY
	fmt.Println("Termino do processo de bufferização do arrquivo ", local_time.Format("03/02/2007 15:04:05"))

	err = os.Remove("lista copy.txt")
	if err != nil {
		panic(err)
	}
}
