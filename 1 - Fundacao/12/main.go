package main

func main() {

	// Memória -> Endereço -> Valor

	// Variável -> ponteiro que tem um endereço na memória -> valor

	a := 10

	var ponteiro *int = &a
	*ponteiro = 20

	b := &a
	println(ponteiro)

	println(*b)

	println(b)

	println(&a)

	println(a)
}
