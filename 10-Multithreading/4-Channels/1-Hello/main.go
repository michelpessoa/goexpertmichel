package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // Vazio

	// Thread 2
	go func() {
		channel <- "Olá Mundo!" // Está cheio
	}()

	// Thread 1
	msg := <-channel // Canal esvazia
	fmt.Println(msg)
}
