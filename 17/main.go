package main

import "fmt"

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
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
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("A soma dos valores do map é %v\n", SomaInteiro(m))

	m2 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	fmt.Printf("A soma dos valores do map é %v\n", SomaFloat(m2))

	m3 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	fmt.Printf("A soma dos valores do map é %v\n", Soma(m3))

	m4 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("A soma dos valores do map é %v\n", Soma(m4))
}
