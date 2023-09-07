package main

import "fmt"

const a = "Olá mundo!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Michel"
	e float64 = 1.2
	f ID      = 2
)

func main() {
	fmt.Printf("O tipo de E é %T", e)
	fmt.Printf("\nO tipo de E é %v", e)
	fmt.Printf("\n\nO tipo de F é %T", f)
}
