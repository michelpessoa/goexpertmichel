package main

type MyNumber int

// Constraint
type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

// Junção das funções abaixo
func SomaJuncao[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {

	m := map[string]int{"Michel": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Michel": 100.20, "João": 200.20, "Maria": 300.10}
	m3 := map[string]MyNumber{"Michel": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	print(Compara(10, 10))

}
