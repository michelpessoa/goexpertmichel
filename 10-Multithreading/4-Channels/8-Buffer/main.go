package main

func main() {

	// Nós conseguimos lançar mais valores em um canal(chan)
	// antes que o mesmo seja consumido, mas devemos ter
	// cautela em utilização deste recurso
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
